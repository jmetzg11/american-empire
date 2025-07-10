package api

import (
	"errors"
	"fmt"
	"good-guys/backend/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetEvents(c *gin.Context) {
	var events []models.Event
	query := h.DB.Select("id, title, date, country")
	query = query.Preload("Tags")
	query = query.Where("active IS NOT NULL")
	result := query.Find(&events)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to get events"})
		return
	}

	var response []models.MainTableResponse
	for _, event := range events {
		var tags []string
		for _, tag := range event.Tags {
			tags = append(tags, tag.Name)
		}
		response = append(response, models.MainTableResponse{
			ID:      event.ID,
			Title:   event.Title,
			Date:    event.Date,
			Country: event.Country,
			Tags:    tags,
		})
	}
	c.JSON(200, response)
}

func (h *Handler) GetEvent(c *gin.Context) {
	var request models.EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	query := h.DB.Preload("Tags")
	query = query.Preload("Sources")
	query = query.Preload("Medias")
	query = query.Where("id = ?", request.ID)
	result := query.First(&event)

	if result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(404, gin.H{"error": "Event not found"})
	}

	c.JSON(200, event)
}

func addTagsToContriubtionEvent(tx *gorm.DB, event *models.Event, tags string) error {
	tagNames := strings.Split(tags, ", ")
	var eventTags []models.Tag

	for _, tagName := range tagNames {
		tagName = strings.TrimSpace(tagName)
		if tagName == "" {
			continue
		}

		var tag models.Tag
		if err := tx.Where("name = ?", tagName).First(&tag).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				tag = models.Tag{Name: tagName}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}
		eventTags = append(eventTags, tag)
	}
	if err := tx.Model(event).Association("Tags").Replace(eventTags); err != nil {
		return err
	}

	return nil
}

func (h *Handler) ContributeEvent(c *gin.Context) {
	var event models.Event

	event.Title = c.PostForm("title")
	event.Country = c.PostForm("country")
	event.Description = c.PostForm("description")

	dateStr := c.PostForm("date")
	parsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	event.Date = parsedDate

	tx := h.DB.Begin()

	if err := h.DB.Create(&event).Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": "Failed to create event"})
		return
	}

	tags := c.PostForm("tags")
	if tags != "" {
		if err := addTagsToContriubtionEvent(tx, &event, tags); err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to add tags"})
			return
		}
	}

	form, _ := c.MultipartForm()

	for i := 0; ; i++ {
		nameKey := fmt.Sprintf("source[%d][name]", i)
		urlKey := fmt.Sprintf("source[%d][url]", i)

		names, nameExists := form.Value[nameKey]
		urls, urlExists := form.Value[urlKey]

		if !nameExists || !urlExists {
			break
		}

		source := models.Source{
			EventID: event.ID,
			Name:    names[0],
			URL:     urls[0],
		}
		if err := tx.Create(&source).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to create source"})
			return
		}
	}

	for i := 0; ; i++ {
		typeKey := fmt.Sprintf("media[%d][type]", i)
		mediaTypes, exists := form.Value[typeKey]

		if !exists {
			break
		}

		media := models.Media{
			EventID: event.ID,
			Type:    mediaTypes[0],
		}

		if mediaTypes[0] == "photo" {
			files := form.File[fmt.Sprintf("media[%d][file]", i)]
			if len(files) > 0 {
				path, err := saveUploadedPhoto(c, files[0], event.ID)
				if err != nil {
					c.JSON(500, gin.H{"error": "Failed to save photo"})
					return
				}
				media.Path = path
			}
			if captions := form.Value[fmt.Sprintf("media[%d][caption]", i)]; len(captions) > 0 {
				media.Caption = captions[0]
			}
		} else if mediaTypes[0] == "youtube" {
			if urls := form.Value[fmt.Sprintf("media[%d][url]", i)]; len(urls) > 0 {
				media.URL = urls[0]
			}
			if captions := form.Value[fmt.Sprintf("media[%d][caption]", i)]; len(captions) > 0 {
				media.Caption = captions[0]
			}
		}
		if err := tx.Create(&media).Error; err != nil {
			tx.Rollback()
			c.JSON(500, gin.H{"error": "Failed to create media"})
			return
		}
	}

	tx.Commit()

	notifyAdmin(form.Value["email"][0], &event)

	c.JSON(200, gin.H{"message": "Event contributed successfully"})
}

func (h *Handler) GetTags(c *gin.Context) {
	var tags []models.Tag
	result := h.DB.Order("LOWER(name) ASC").Find(&tags)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to get tags"})
		return
	}

	c.JSON(200, tags)
}

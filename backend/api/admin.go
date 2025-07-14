package api

import (
	"american-empire/backend/database"
	"american-empire/backend/models"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) GetAdminEvents(c *gin.Context) {
	var events []models.Event
	h.DB.Select("id, title, date, country").Where("active IS NULL").Find(&events)

	var response []models.MainTableResponse
	for _, event := range events {
		response = append(response, models.MainTableResponse{
			ID:      event.ID,
			Title:   event.Title,
			Date:    event.Date,
			Country: event.Country,
		})
	}
	c.JSON(200, response)
}

func updateEventTags(tx *gorm.DB, event *models.Event, tagsStr string) error {
	tagNames := strings.Split(tagsStr, ", ")
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

	// clean up orphaned tags
	tx.Where("id NOT IN (SELECT DISTINCT tag_id FROM event_tags)").Delete(&models.Tag{})

	return nil
}

func (h *Handler) EditEvent(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	eventID, exists := payload["id"]
	if !exists {
		c.JSON(400, gin.H{"error": "Event ID is required"})
		return
	}

	var event models.Event
	query := h.DB.Preload("Sources")
	query = query.Preload("Medias")
	result := query.First(&event, eventID)
	if result.Error != nil {
		log.Println("Failed to get event", result.Error)
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	tx := h.DB.Begin()

	if title, ok := payload["Title"].(string); ok {
		event.Title = title
	}
	if desc, ok := payload["Description"].(string); ok {
		event.Description = desc
	}
	if country, ok := payload["Country"].(string); ok {
		event.Country = country
	}
	if dateStr, ok := payload["Date"].(string); ok {
		if date, err := time.Parse("2006-01-02", dateStr); err == nil {
			event.Date = date
		}
	}

	if err := tx.Save(&event).Error; err != nil {
		tx.Rollback()
		log.Println("Failed to update event", err)
		c.JSON(500, gin.H{"error": "Failed to update event"})
		return
	}

	if tags, ok := payload["Tags"].(string); ok {
		if err := updateEventTags(tx, &event, tags); err != nil {
			tx.Rollback()
			log.Println("Failed to update event tags", err)
			c.JSON(500, gin.H{"error": "Failed to update event tags"})
			return
		}
	}

	for key, value := range payload {
		if strings.HasPrefix(key, "media-") {
			mediaIDStr := strings.TrimPrefix(key, "media-")
			mediaID, err := strconv.ParseUint(mediaIDStr, 10, 32)
			if err != nil {
				continue
			}

			if mediaData, ok := value.(map[string]interface{}); ok {
				if caption, ok := mediaData["Caption"].(string); ok {
					tx.Model(&models.Media{}).Where("id = ?", mediaID).Update("caption", caption)
				}
			}
		}

		if strings.HasPrefix(key, "source-") {
			sourceIDStr := strings.TrimPrefix(key, "source-")
			sourceID, err := strconv.ParseUint(sourceIDStr, 10, 32)
			if err != nil {
				continue
			}

			if sourceData, ok := value.(map[string]interface{}); ok {
				if name, ok := sourceData["Name"].(string); ok {
					tx.Model(&models.Source{}).Where("id = ?", sourceID).Update("name", name)
				}
			}
		}
	}

	tx.Commit()
	c.JSON(200, gin.H{"message": "Event edited"})
}

func (h *Handler) ApproveEvent(c *gin.Context) {
	var request models.EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	now := time.Now()
	result := h.DB.Model(&models.Event{}).Where("id = ?", request.ID).Update("active", &now)

	if result.Error != nil {
		log.Println("Failed to approve event", result.Error)
		c.JSON(500, gin.H{"error": "Failed to approve event"})
		return
	}

	if result.RowsAffected == 0 {
		log.Println("Event not found")
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Event approved"})
}

func (h *Handler) UnapproveEvent(c *gin.Context) {
	var request models.EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	result := h.DB.Model(&models.Event{}).Where("id = ?", request.ID).Update("active", nil)

	if result.Error != nil {
		log.Println("Failed to unapprove event", result.Error)
		c.JSON(500, gin.H{"error": "Failed to unapprove event"})
		return
	}

	if result.RowsAffected == 0 {
		log.Println("Event not found")
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(200, gin.H{"message": "Event unapproved"})
}

func getMediaData(c *gin.Context) (models.Media, error) {
	eventId, err := strconv.ParseUint(c.PostForm("event_id"), 10, 32)
	if err != nil {
		return models.Media{}, err
	}

	return models.Media{
		EventID: uint(eventId),
		Caption: c.PostForm("title"),
	}, nil
}

func (h *Handler) UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Failed to get media data", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	media, err := getMediaData(c)
	if err != nil {
		log.Println("Failed to get media data", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	media.Type = "photo"

	path, err := saveUploadedPhoto(c, file, media.EventID)
	if err != nil {
		log.Println("Failed to save photo", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	media.Path = path
	h.DB.Create(&media)

	c.JSON(200, gin.H{"message": "Photo uploaded"})
}

func (h *Handler) UploadYoutube(c *gin.Context) {
	media, err := getMediaData(c)
	if err != nil {
		log.Println("Failed to get media data", err)
		c.JSON(400, gin.H{"error": err.Error()})
	}

	media.Type = "youtube"
	media.URL = c.PostForm("url")

	h.DB.Create(&media)
	c.JSON(200, gin.H{"message": "Youtube uploaded"})
}

func (h *Handler) DeleteMedia(c *gin.Context) {
	var request models.MediaDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var media models.Media
	result := h.DB.Where("id = ?", request.ID).First(&media)

	if result.Error != nil {
		log.Println("Failed to get media", result.Error)
		c.JSON(404, gin.H{"error": "Media not found"})
		return
	}

	if media.Path != "" {
		if os.Getenv("GIN_MODE") == "release" {
			database.SupabaseClient.Storage.RemoveFile("photos", []string{media.Path})
		} else {
			fullPath := fmt.Sprintf("data/photos/%s", media.Path)
			os.Remove(fullPath)
		}
	}

	h.DB.Delete(&media)
	c.JSON(200, gin.H{"message": "Media deleted"})
}

func (h *Handler) DeleteSources(c *gin.Context) {
	var request models.SourceDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var source models.Source
	result := h.DB.Where("id = ?", request.ID).First(&source)

	if result.Error != nil {
		log.Println("Failed to get source", result.Error)
		c.JSON(404, gin.H{"error": "Source not found"})
		return
	}

	h.DB.Delete(&source)
	c.JSON(200, gin.H{"message": "Media deleted"})
}

func (h *Handler) AddSources(c *gin.Context) {
	var request models.SourceAddRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Failed to bind JSON", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	eventID, err := strconv.ParseUint(request.EventID, 10, 32)
	if err != nil {
		log.Println("Failed to parse event ID", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	source := models.Source{
		EventID: uint(eventID),
		Name:    request.Name,
		URL:     request.URL,
	}

	if err := h.DB.Create(&source).Error; err != nil {
		log.Println("Failed to create source", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Source added"})
}

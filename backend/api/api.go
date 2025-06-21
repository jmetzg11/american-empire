package api

import (
	"fmt"
	"good-guys/backend/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetEvents(c *gin.Context) {
	var events []models.Event
	h.DB.Select("id, title, date, country").Find(&events)

	var response []models.DataResponse
	for _, event := range events {
		response = append(response, models.DataResponse{
			ID:      event.ID,
			Title:   event.Title,
			Date:    event.Date,
			Country: event.Country,
		})
	}
	c.JSON(200, response)
}

func (h *Handler) GetEvent(c *gin.Context) {
	var request models.EventRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	query := h.DB.Select("id, title, date, description")
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

	if err := h.DB.Create(&event).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create event"})
		return
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
		h.DB.Create(&source)
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
				file := files[0]
				filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
				dir := fmt.Sprintf("data/photos/%d", event.ID)
				os.MkdirAll(dir, 0755)
				fullPath := fmt.Sprintf("%s/%s", dir, filename)

				if err := c.SaveUploadedFile(file, fullPath); err != nil {
					continue
				}
				media.Path = fmt.Sprintf("%d/%s", event.ID, filename)
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

		h.DB.Create(&media)
	}

	c.JSON(200, gin.H{"message": "Event contributed successfully"})
}

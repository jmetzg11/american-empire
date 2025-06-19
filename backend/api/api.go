package api

import (
	"fmt"
	"good-guys/backend/models"
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
	parrsedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}
	event.Date = parrsedDate

	if err := h.DB.Create(&event).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create event"})
		return
	}

	form, _ := c.MultipartForm()

	if sourceNames, exists := form.Value["sources[][name]"]; exists {
		sourceUrls := form.Value["sources[][url]"]
		for i, name := range sourceNames {
			if i < len(sourceUrls) {
				source := models.Source{
					EventID: event.ID,
					Name:    name,
					URL:     sourceUrls[i],
				}
				h.DB.Create(&source)
			}
		}
	}

	// if mediaTypes, exists := form.Value["media[][type]"]; exists {
	// 	for i, mediaType := range mediaTypes {
	// 		media := models.Media{
	// 			EventID: event.ID,
	// 			Type:    mediaType,
	// 		}

	// 		if mediaType == "photo" {
	// 			files := form.File[fmt.Sprintf("media[%d][file]", i)]
	// 			if len(files) > 0 {
	// 				file := files[0]
	// 				filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	// 				filepath := fmt.Sprintf("uploads/%s", filename)

	// 				if err := c.SaveUploadedFile(file, filepath); err != nil {
	// 					continue
	// 				}
	// 				media.Path = filepath
	// 			}
	// 			if captions := form.Value[fmt.Sprintf("media[%d][caption]", i)]; len(captions) > 0 {
	// 				media.Caption = captions[0]
	// 			}
	// 		} else if mediaType == "youtube" {
	// 			if urls := form.Value[fmt.Sprintf("media[%d][url]", i)]; len(urls) > 0 {
	// 				media.URL = urls[0]
	// 			}
	// 			if captions := form.Value[fmt.Sprintf("media[%d][caption]", i)]; len(captions) > 0 {
	// 				media.Caption = captions[0]
	// 			}
	// 		}

	// 		h.DB.Create(&media)
	// 	}
	// }

	c.JSON(200, gin.H{"message": "Event contributed successfully"})
}

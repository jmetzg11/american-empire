package api

import (
	"fmt"
	"good-guys/backend/models"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAdminEvents(c *gin.Context) {
	var events []models.Event
	h.DB.Select("id, title, date, country").Where("active IS NULL").Find(&events)

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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	media, err := getMediaData(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	media.Type = "photo"

	path, err := saveUploadedPhoto(c, file, media.EventID)
	if err != nil {
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
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var media models.Media
	result := h.DB.Where("id = ?", request.ID).First(&media)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Media not found"})
		return
	}

	if media.Path != "" {
		fullPath := fmt.Sprintf("data/photos/%s", media.Path)
		os.Remove(fullPath)
	}

	h.DB.Delete(&media)
	c.JSON(200, gin.H{"message": "Media deleted"})
}

func (h *Handler) DeleteSources(c *gin.Context) {
	var request models.SourceDeleteRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var source models.Source
	result := h.DB.Where("id = ?", request.ID).First(&source)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Source not found"})
		return
	}

	h.DB.Delete(&source)
	c.JSON(200, gin.H{"message": "Media deleted"})
}

func (h *Handler) AddSources(c *gin.Context) {
	var request models.SourceAddRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	eventID, err := strconv.ParseUint(request.EventID, 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	source := models.Source{
		EventID: uint(eventID),
		Name:    request.Name,
		URL:     request.URL,
	}

	if err := h.DB.Create(&source).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Source added"})
}

package api

import (
	"fmt"
	"good-guys/backend/models"

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

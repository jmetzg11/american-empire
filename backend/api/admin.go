package api

import (
	"good-guys/backend/models"

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

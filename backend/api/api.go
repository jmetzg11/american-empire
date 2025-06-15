package api

import (
	"good-guys/backend/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetData(c *gin.Context) {
	var events []models.Event
	h.DB.Find(&events)
	c.JSON(200, events)
}

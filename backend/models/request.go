package models

type EventRequest struct {
	ID string `json:"id" binding:"required"`
}

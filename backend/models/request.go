package models

type EventRequest struct {
	ID string `json:"id" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type MediaDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type SourceDeleteRequest struct {
	ID uint `json:"id" binding:"required"`
}

type SourceAddRequest struct {
	EventID string `json:"EventID" binding:"required"`
	Name    string `json:"Name" binding:"required"`
	URL     string `json:"URL" binding:"required"`
}

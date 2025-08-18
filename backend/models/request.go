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


type NewBookRequest struct {
	Title string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Link string `json:"link" binding:"required"`
	Events string `json:"events" binding:"required"`
	SelectedTags []uint `json:"selectedTags" binding:"required"`
}

package models

import "time"

type DataResponse struct {
	ID      uint      `json:"id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Country string    `json:"country"`
}

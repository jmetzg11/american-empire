package models

import "time"

type Event struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"not null"`
	Date        time.Time `gorm:"not null;index"`
	Country     string    `gorm:"not null;index"`
	Description string    `gorm:"type:text;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Active      *time.Time
	Flagged     bool `gorm:"default:false"`

	Sources []Source `gorm:"foreignKey:EventID;"`
	Medias  []Media  `gorm:"foreignKey:EventID"`
}

type Source struct {
	ID      uint   `gorm:"primaryKey"`
	EventID uint   `gorm:"not null"`
	Name    string `gorm:"not null"`
	URL     string `gorm:"not null"`
}

type Media struct {
	ID      uint `gorm:"primaryKey"`
	EventID uint `gorm:"not null"`
	Type    string
	URL     string
	Path    string
}

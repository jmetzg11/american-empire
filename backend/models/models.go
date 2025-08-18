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
	Tags    []Tag    `gorm:"many2many:event_tags;"`
	Books []Book `gorm:"many2many:book_events;"`
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
	Caption string
}

type Tag struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"not null;unique"`
	Events []Event `gorm:"many2many:event_tags;"`
}

type Book struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"not null"`
	Author string `gorm:"not null"`
	Link   string `gorm:"not null"`

	Events []Event `gorm:"many2many:book_events;"`
	Tags   []Tag   `gorm:"many2many:book_tags;"`
}

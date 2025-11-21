package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Completed   bool   `gorm:"not null"`
	UserID      string
	AjoloteID   string
}

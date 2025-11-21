package models

import (
	models "Gin/src/domain/models/task"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string        `gorm:"not null"`
	Email    string        `gorm:"not null;unique"`
	Password string        `gorm:"not null"`
	Tasks    []models.Task `gorm:"foreignKey:UserID;references:ID"`
}

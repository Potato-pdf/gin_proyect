package models

import (
	models "Gin/src/domain/models/task"

	"gorm.io/gorm"
)

type Ajolote struct {
	gorm.Model
	Id     string        `gorm:"primaryKey;not null"`
	Nombre string        `gorm:"not null"`
	Color  string        `gorm:"not null"`
	Task   []models.Task `gorm:"foreignKey:AjoloteID;references:Id"`
}

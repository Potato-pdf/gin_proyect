package models

import (
	"Gin/src/domain/models/task"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string      `json:"name"`
	Email    string      `json:"email" gorm:"unique"`
	Password string      `json:"-"`
	Tasks    []task.Task `json:"tasks"`
}

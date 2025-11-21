package models

import "gorm.io/gorm"

type Ajolote struct {
	gorm.Model
	Nombre string
	Color  string
}

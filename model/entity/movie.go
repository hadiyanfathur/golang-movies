package entity

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string
	Description string
}

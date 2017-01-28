package models

import (
	"github.com/jinzhu/gorm"
)

type Recipe struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	Description string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
}

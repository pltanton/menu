package models

import (
	"github.com/jinzhu/gorm"
)

type Recipe struct {
	gorm.Model
	Name        string `gorm:"primary_key"`
	Description string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
}

package models

import (
	"github.com/jinzhu/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string `gorm:"primary_key"`
}

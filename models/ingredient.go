package models

import (
	"github.com/jinzhu/gorm"
)

type Ingredient struct {
	gorm.Model
	Name string `gorm:"not null;unique"`
}

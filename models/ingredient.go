package models

type Ingredient struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null;unique"`
}

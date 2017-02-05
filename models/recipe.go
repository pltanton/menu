package models

type Recipe struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"not null;unique"`
	Description string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients"`
}

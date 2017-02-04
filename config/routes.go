package config

import (
	"net/http"

	"kaliwe.ru/menu/handlers"
)

func initRoutes() {
	// Ingredient routes
	http.HandleFunc("/ingredients", handlers.Ingredients)
	http.HandleFunc("/ingredient/new", handlers.IngredientNew)
}

package config

import (
	"github.com/gin-gonic/gin"
	"kaliwe.ru/menu/handlers"
)

func initRoutes(router *gin.Engine) {
	// Ingredient routes
	router.GET("/", handlers.Ingredients)
	router.GET("/ingredients", handlers.Ingredients)
	router.POST("/ingredient/new", handlers.IngredientNew)
}

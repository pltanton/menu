package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/handlers"
)

func main() {
	// Ensure to generate instance of db on first start
	db.GetInstance()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			nil,
		)
	})

	// Ingredient routes
	router.GET("/ingredient/new", handlers.ShowNewIngredient)
	router.POST("/ingredient/new", handlers.NewIngredient)

	router.Run()
}

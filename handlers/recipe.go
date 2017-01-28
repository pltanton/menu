package handlers

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

func RecipeNew(c *gin.Context) {
	db := db.GetInstance()
	name, _ := c.GetPostForm("name")
	description, _ := c.GetPostForm("description")
	recipe := models.Recipe{Name: name, Description: description}

	db = db.Create(&recipe)

	if err := db.Error; err != nil {
		c.Redirect(http.StatusBadRequest, "/recipe/new")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/recipes")
	}
}

func Recipes(c *gin.Context) {
	recipes := []models.Recipe{}
	db.GetInstance().Find(&recipes)
	c.HTML(
		http.StatusOK,
		"recipes.html",
		gin.H{
			"recipes": recipes,
		},
	)
}

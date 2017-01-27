package handlers

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

func Ingredient(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"ingredient_new.html",
		nil,
	)
}

func IngredientNew(c *gin.Context) {
	db := db.GetInstance()
	name, _ := c.GetPostForm("name")
	ingredient := models.Ingredient{Name: name}

	db = db.Create(&ingredient)

	if err := db.Error; err != nil {
		c.Redirect(http.StatusBadRequest, "/ingredient/new")
	} else {
		c.Redirect(http.StatusMovedPermanently, "/ingredients")
	}
}

func Ingredients(c *gin.Context) {
	ingredients := []models.Ingredient{}
	db.GetInstance().Find(&ingredients)
	c.HTML(
		http.StatusOK,
		"ingredients.html",
		gin.H{
			"ingredients": ingredients,
		},
	)

}

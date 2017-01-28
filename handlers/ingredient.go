package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

func IngredientNew(c *gin.Context) {
	db := db.GetInstance()
	name := c.Query("name")
	fmt.Println(name)
	fmt.Println("mamke ebal")
	ingredient := models.Ingredient{Name: name}
	db = db.Create(&ingredient)

	if err := db.Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, nil)
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

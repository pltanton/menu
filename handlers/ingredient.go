package handlers

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

func ShowNewIngredient(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"ingredient_new.html",
		nil,
	)
}

func NewIngredient(c *gin.Context) {
	name, _ := c.GetPostForm("name")
	ingredient := models.Ingredient{Name: name}
	db.GetInstance().Create(&ingredient)
}

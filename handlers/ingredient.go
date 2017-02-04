package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

func IngredientNew(w http.ResponseWriter, r *http.Request) {
	db := db.GetInstance()
	name := r.URL.Query().Get("name")
	ingredient := models.Ingredient{Name: name}
	db = db.Create(&ingredient)

	if err := db.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(err)
	}
}

func Ingredients(w http.ResponseWriter, r *http.Request) {
	ingredients := []models.Ingredient{}
	db.GetInstance().Find(&ingredients)
	_, err := RespondAce(w, "ingredients", map[string]interface{}{
		"Ingredients": ingredients,
	})
	if err != nil {
		fmt.Println(err)
	}
}

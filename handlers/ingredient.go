package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"kaliwe.ru/menu/db"
	"kaliwe.ru/menu/models"
)

// IngredientNew creates new ingredient by given name
func IngredientNew(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := db.GetInstance()
	name := vars["name"]
	ingredient := models.Ingredient{Name: name}
	db = db.Create(&ingredient)

	if err := db.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
}

// IngredientDelete soft deletes ingredient by given id in url
func IngredientDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := db.GetInstance()
	ingredientID, _ := strconv.Atoi(vars["id"])
	db.Delete(models.Ingredient{}, ingredientID)
	if err := db.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}
}

// Ingredient provides a view for single ingredient
func Ingredient(w http.ResponseWriter, r *http.Request) {

}

// Ingredients provides a view for all ingredients
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

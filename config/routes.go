package config

import (
	"net/http"

	"github.com/gorilla/mux"

	"kaliwe.ru/menu/handlers"
)

// Route describes structure of a single route with single method for CRUD
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Routes is alias for array of Route
type Routes []Route

func initRoutes() {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	http.Handle("/", router)
}

var routes = Routes{
	Route{
		"IngredientsShow",
		"GET",
		"/ingredients",
		handlers.Ingredients,
	},
	Route{
		"IngredientNew",
		"POST",
		"/ingredient/new",
		handlers.IngredientNew,
	},
	Route{
		"IngredientDelete",
		"DELETE",
		"/ingredient/{id:[0-9]+}",
		handlers.IngredientDelete,
	},
}

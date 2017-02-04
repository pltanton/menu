package config

import (
	"net/http"
)

// Init initializes application
func Init() {

	initDB()
	initRoutes()
	initStatics()

	http.ListenAndServe(":8080", nil)
}

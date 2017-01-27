package config

import (
	"github.com/gin-gonic/gin"
)

// Init initializes application
func Init() *gin.Engine {
	// TODO: move to another package
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// TODO: move to another package
	router.Static("/javascripts", "assets")
	router.Static("/stylesheets", "assets")

	initRoutes(router)
	initDB()

	return router
}

package config

import (
	"github.com/gin-gonic/gin"
)

// Init initializes application
func Init() *gin.Engine {
	router := gin.Default()

	// TODO: move to another package
	router.LoadHTMLGlob("templates/*")

	initStatics(router)
	initRoutes(router)
	initDB()

	return router
}

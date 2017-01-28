package config

import (
	"github.com/gin-gonic/gin"
)

func initStatics(router *gin.Engine) {
	router.Static("/assets", "assets")
	router.Static("/vendor", "vendor/assets")
}

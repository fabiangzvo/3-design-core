package api

import (
	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func noRoute(c *gin.Context) {
	c.JSON(404, gin.H{"message": "Service not found"})
}

//Router load all available routes
func Router(router *gin.Engine) {
	router.GET("/health-check", healthCheck)
	router.NoRoute(noRoute)
}

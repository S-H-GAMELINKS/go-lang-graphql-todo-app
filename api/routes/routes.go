package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Router) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
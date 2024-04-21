package main

import (
	"go-crud-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
		initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // changed port to 3000 in .env file
}
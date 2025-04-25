package main

import (
	"go-repl/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main(){
	router := gin.Default()

	// Welcome route
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go Playground REPL API 🚀",
		})
	})

	// Health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Set up the routes and attach the handlers
	router.POST("/run", handler.RunCode)

	//start the server on port 8080
	router.Run(":8080")

	

}

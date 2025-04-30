package main

import (
	"go-repl/internal/handler"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin router
	router := gin.Default()

	// Configure CORS to allow your frontend domain
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://goplay-mocha.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go Playground REPL API 🚀",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Main code execution endpoint
	router.POST("/run", handler.RunCode)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		panic("failed to start server: " + err.Error())
	}
}

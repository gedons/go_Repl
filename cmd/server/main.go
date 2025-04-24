package main

import (  
  "github.com/gin-gonic/gin"
  "go-repl/internal/handler"
)


func main(){
	router := gin.Default()

	// Set up the routes and attach the handlers
	router.POST("/run", handler.RunCode)

	//start the server on port 8080
	router.Run(":8080")
}

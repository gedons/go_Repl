package handler

import (
	"net/http"
	"go-repl/internal/runner"

	"github.com/gin-gonic/gin"
)

type RunRequest struct {
	Code string `json:"code" binding:"required"`
}

type RunResponse struct {
	Output string `json:"output"`
	Error  string `json:"error"`
}

func RunCode(c *gin.Context) {
	var req RunRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Execute the code using the runner package
	// This function will handle the logic of running the code in a Docker container
	output, execErr := runner.ExecuteCode(req.Code)

	c.JSON(http.StatusOK, RunResponse{
		Output: output,
		Error:  execErr,
	})
}
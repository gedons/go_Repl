package handler

import (
	"go-repl/internal/runner"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RunRequest struct {
	Code  string `json:"code"`
	Stdin string `json:"stdin"`
}

func RunCode(c *gin.Context) {
	var req RunRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	stdout, stderr := runner.ExecuteCode(req.Code, req.Stdin)

	c.JSON(http.StatusOK, gin.H{
		"stdout": stdout,
		"stderr": stderr,
	})
}

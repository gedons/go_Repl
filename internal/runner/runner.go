package runner

import (
	"bytes"
	"fmt"
	"os/exec"
	"go-repl/internal/utils"
)

func ExecuteCode(code string) (string, string) {
	// Step 1: Save code to temp.go
	filePath, err := utils.SaveCodeToFile(code)
	if err != nil {
		return "", fmt.Sprintf("Failed to save code: %v", err)
	}

	// Step 2: Run Docker command
	cmd := exec.Command("docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app/temp.go", filePath),
		"-w", "/app",
		"golang:1.21-alpine", 
		"go", "run", "temp.go",
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String()
	}

	return stdout.String(), ""
}

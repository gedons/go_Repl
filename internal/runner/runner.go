package runner

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
	"go-repl/internal/utils"
)


func logExecutionStep(step string) {
	log.Printf("[INFO]: %s", step)
}

func cleanUpTempFiles(dir string) error {
	err := os.RemoveAll(dir)
	if err != nil {
		return fmt.Errorf("failed to clean up temp files: %v", err)
	}
	logExecutionStep(fmt.Sprintf("Cleaned up temp files from %s", dir))
	return nil
}

func ExecuteCode(code string) (string, string) {

	// Log the start of execution
	logExecutionStep("Starting code execution")

	// Save code to temp.go
	_, dir, err := utils.SaveCodeToFile(code)
	if err != nil {
		logExecutionStep(fmt.Sprintf("Failed to save code: %v", err))
		return "", fmt.Sprintf("Failed to save code: %v", err)
	}
	logExecutionStep("Code saved to file")


	// Set up timeout context (5 seconds)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	// Run Docker command with context
	cmd := exec.CommandContext(ctx, "docker", "run", "--rm",
		"-v", fmt.Sprintf("%s:/app", dir),
		"-w", "/app",
		"golang:1.21-alpine",
		"go", "run", "temp.go",
	)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		logExecutionStep("Error executing command")
		if ctx.Err() == context.DeadlineExceeded {
			logExecutionStep("Execution timed out")
			return "", "Code execution timed out"
		}
		logExecutionStep(fmt.Sprintf("Execution failed: %v", stderr.String()))
		return stdout.String(), stderr.String()
	}

	logExecutionStep("Execution successful")

	// Clean up temp files after execution	
	err = cleanUpTempFiles(dir)
	if err != nil {
		logExecutionStep(fmt.Sprintf("Error cleaning up temp files: %v", err))
	}


	return stdout.String(), ""
}


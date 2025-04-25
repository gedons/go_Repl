package runner

import (
	"bytes"
	"fmt"
	"go-repl/internal/utils"
	"log"
	"os"
	"os/exec"
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

    // Log the command that will be executed
    logExecutionStep(fmt.Sprintf("Running Docker command: docker run --rm -v %s:/app -w /app golang:1.21-alpine go run temp.go", dir))

    // Run Docker command to execute the code
    cmd := exec.Command("docker", "run", "--rm",
        "-v", fmt.Sprintf("%s:/app", dir),
        "-w", "/app",
        "golang:1.21-alpine",
        "go", "run", "temp.go",
    )

    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Execute the command
    err = cmd.Run()
    if err != nil {
        logExecutionStep("Error executing command")
        logExecutionStep(fmt.Sprintf("Execution failed: %v", stderr.String()))
        return stdout.String(), stderr.String()
    }

    logExecutionStep("Execution successful")

    // Clean up temp files after execution
    err = cleanUpTempFiles(dir)
    if err != nil {
        logExecutionStep(fmt.Sprintf("Error cleaning up temp files: %v", err))
    }

    // Return the standard output or error message
    return stdout.String(), ""
}

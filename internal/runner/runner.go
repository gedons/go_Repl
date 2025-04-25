package runner

import (
	"fmt"
	"go-repl/internal/utils"
	"log"
	"os"
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

    // Now we don't need to run Docker command, just return success
    logExecutionStep("EBS will use prebuilt image from ECR")

    // Clean up temp files after execution
    err = cleanUpTempFiles(dir)
    if err != nil {
        logExecutionStep(fmt.Sprintf("Error cleaning up temp files: %v", err))
    }

    return "Code execution successfully.", ""
}



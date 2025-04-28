// internal/runner/runner.go
package runner

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"go-repl/internal/utils"
)

func logStep(msg string) {
	log.Printf("[RUNNER] %s\n", msg)
}

// ExecuteCode compiles and runs user code, returning its stdout or stderr.
func ExecuteCode(code string) (string, string) {
	// 1) Save user code
	srcPath, dir, err := utils.SaveCodeToFile(code)
	if err != nil {
		logStep(fmt.Sprintf("Save error: %v", err))
		return "", fmt.Sprintf("save error: %v", err)
	}
	defer os.RemoveAll(dir)
	logStep("Code written to " + srcPath)

	// 2) Compile temp.go -> tempbin
	binPath := filepath.Join(dir, "tempbin")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	logStep("Compiling user code...")
	build := exec.CommandContext(ctx, "go", "build", "-o", binPath, srcPath)
	var buildOut, buildErr bytes.Buffer
	build.Stdout = &buildOut
	build.Stderr = &buildErr
	err = build.Run()

	if err != nil {
		logStep(fmt.Sprintf("Compilation failed: %v", err))
		logStep(fmt.Sprintf("Compilation stderr: %s", buildErr.String()))
		return "", fmt.Sprintf("compile error: %v\nstderr: %s", err, buildErr.String())
	}

	logStep("Compilation successful")

	// 3) Run compiled binary
	logStep("Running user binary...")
	run := exec.CommandContext(ctx, binPath)
	var out, runErr bytes.Buffer
	run.Stdout = &out
	run.Stderr = &runErr

	err = run.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return out.String(), "execution timed out"
		}
		logStep(fmt.Sprintf("Run error: %v", err))
		return out.String(), runErr.String()
	}

	logStep("Execution finished successfully")
	return out.String(), ""
}

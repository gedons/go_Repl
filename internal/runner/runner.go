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
		return "", fmt.Sprintf("save error: %v", err)
	}
	defer os.RemoveAll(dir)
	logStep("code written to " + srcPath)

	// 2) Compile temp.go -> tempbin
	binPath := filepath.Join(dir, "tempbin")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	logStep("compiling user code")
	build := exec.CommandContext(ctx, "go", "build", "-o", binPath, srcPath)
	var buildErr bytes.Buffer
	build.Stderr = &buildErr
	if err := build.Run(); err != nil {
		return "", buildErr.String()
	}

	// 3) Run the compiled binary
	logStep("running user binary")
	run := exec.CommandContext(ctx, binPath)
	var out, runErr bytes.Buffer
	run.Stdout = &out
	run.Stderr = &runErr

	if err := run.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return out.String(), "execution timed out"
		}
		return out.String(), runErr.String()
	}

	logStep("execution finished successfully")
	return out.String(), ""
}

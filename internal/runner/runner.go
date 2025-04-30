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
func ExecuteCode(code string, stdin string) (string, string) {
	// 1) Save user code
	srcPath, dir, err := utils.SaveCodeToFile(code)
	if err != nil {
		logStep(fmt.Sprintf("Save error: %v", err))
		return "", fmt.Sprintf("save error: %v", err)
	}
	defer os.RemoveAll(dir)
	logStep("Code written to " + srcPath)

	// 2) Compile the code
	binPath := filepath.Join(dir, "tempbin")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	build := exec.CommandContext(ctx, "go", "build", "-o", binPath, srcPath)
	var buildOut, buildErr bytes.Buffer
	build.Stdout = &buildOut
	build.Stderr = &buildErr
	err = build.Run()
	if err != nil {
		return "", fmt.Sprintf("compile error: %v\nstderr: %s", err, buildErr.String())
	}

	// 3) Run the compiled binary with optional stdin
	run := exec.CommandContext(ctx, binPath)
	if stdin != "" {
		run.Stdin = bytes.NewBufferString(stdin)
	}

	var out, runErr bytes.Buffer
	run.Stdout = &out
	run.Stderr = &runErr

	err = run.Run()
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return out.String(), "execution timed out"
		}
		return out.String(), runErr.String()
	}

	return out.String(), ""
}


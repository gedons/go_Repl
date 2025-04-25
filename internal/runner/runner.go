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

func logStep(step string) {
    log.Printf("[REPL] %s\n", step)
}

func cleanup(dir string) {
    if err := os.RemoveAll(dir); err != nil {
        logStep(fmt.Sprintf("cleanup failed: %v", err))
    } else {
        logStep("temp files cleaned")
    }
}

// ExecuteCode compiles temp.go into a binary and runs it under a timeout.
func ExecuteCode(code string) (stdoutStr, stderrStr string) {
    logStep("starting execution")

    // 1) write code
    srcPath, dir, err := utils.SaveCodeToFile(code)
    if err != nil {
        return "", fmt.Sprintf("save error: %v", err)
    }

    // ensure cleanup
    defer cleanup(dir)

    // 2) compile
    binPath := filepath.Join(dir, "tempbin")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    logStep("compiling code")
    build := exec.CommandContext(ctx, "go", "build", "-o", binPath, srcPath)
    var buildOut, buildErr bytes.Buffer
    build.Stdout = &buildOut
    build.Stderr = &buildErr
    if err := build.Run(); err != nil {
        return buildOut.String(), buildErr.String()
    }

    // 3) run
    logStep("running binary")
    run := exec.CommandContext(ctx, binPath)
    var runOut, runErr bytes.Buffer
    run.Stdout = &runOut
    run.Stderr = &runErr
    if err := run.Run(); err != nil {
        // timeout?
        if ctx.Err() == context.DeadlineExceeded {
            return "", "execution timed out"
        }
        return runOut.String(), runErr.String()
    }

    logStep("execution successful")
    return runOut.String(), ""
}

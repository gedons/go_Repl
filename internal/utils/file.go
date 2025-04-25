// internal/utils/file.go
package utils

import (
    "os"
)

func SaveCodeToFile(code string) (filePath string, dir string, err error) {
    dir = "/tmp/repl"
    filePath = dir + "/temp.go"

    if err = os.MkdirAll(dir, 0755); err != nil {
        return "", "", err
    }

    f, err := os.Create(filePath)
    if err != nil {
        return "", "", err
    }
    defer f.Close()

    if _, err = f.WriteString(code); err != nil {
        return "", "", err
    }
    // ensure it’s flushed
    if err = f.Sync(); err != nil {
        return "", "", err
    }

    return filePath, dir, nil
}

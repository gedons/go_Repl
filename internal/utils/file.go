package utils

import (
    "os"
    "path/filepath"
)

func SaveCodeToFile(code string) (string, string, error) {
    dir, err := filepath.Abs("./code-temp") 
    if err != nil {
        return "", "", err
    }

    filePath := filepath.Join(dir, "temp.go")

    err = os.MkdirAll(dir, 0755)
    if err != nil {
        return "", "", err
    }

    file, err := os.Create(filePath)
    if err != nil {
        return "", "", err
    }
    defer file.Close()

    _, err = file.WriteString(code)
    if err != nil {
        return "", "", err
    }

    // Ensure contents are flushed to disk
    if err := file.Sync(); err != nil {
        return "", "", err
    }

    return filePath, dir, nil
}



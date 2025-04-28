package utils

import (
	"os"
	"path/filepath"
)

// SaveCodeToFile saves code to a temporary file and returns file path and dir
func SaveCodeToFile(code string) (string, string, error) {
	dir, err := os.MkdirTemp("", "gorepl")
	if err != nil {
		return "", "", err
	}
	filePath := filepath.Join(dir, "temp.go")
	err = os.WriteFile(filePath, []byte(code), 0644)
	if err != nil {
		return "", "", err
	}
	return filePath, dir, nil
}

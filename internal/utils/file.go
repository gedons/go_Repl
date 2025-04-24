package utils

import (
	"os"
)

func SaveCodeToFile(code string) (string, error) {
	filePath := "/tmp/temp.go"

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.WriteString(code)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

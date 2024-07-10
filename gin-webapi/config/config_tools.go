package config

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// LoadConfig opens the config file and returns the bytes
func LoadConfig(fileName string) ([]byte, error) {
	path, err := os.Getwd()
	ErrorLog(err, true, "get current path error")

	filePath := filepath.Join(path, "config/", fileName)

	file, err := os.Open(filePath)
	ErrorLog(err, true, "open config.json error")

	defer file.Close()

	bytes, err := io.ReadAll(file)
	ErrorLog(err, true, "read config.json error")

	return bytes, nil
}

// ErrorLog logs the error message
func ErrorLog(err error, flag bool, msg string) {
	if err != nil {
		if len(msg) > 0 {
			log.Printf("error: %v, %s\n", err, msg)
		} else {
			log.Printf("error: %v\n", err)
		}

		if flag {
			os.Exit(1)
		}
	}
}

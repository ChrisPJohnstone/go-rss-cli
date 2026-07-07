package internal

import (
	"log"
	"os"
)

func ensureDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if Verbose {
			log.Printf("Dir %s not found. Creating", path)
		}
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

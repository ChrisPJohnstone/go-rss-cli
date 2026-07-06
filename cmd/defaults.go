package cmd

import (
	"os"

	"path/filepath"
)

func defaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".rss-cli"), nil
}

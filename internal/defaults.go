package internal

import (
	"os"

	"path/filepath"
)

func DefaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".rss-cli"), nil
}

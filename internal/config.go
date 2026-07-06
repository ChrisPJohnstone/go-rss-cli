package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Feeds []string `yaml:"feeds"`
}

func LoadConfig(cmd *cobra.Command, args []string) error {
	path := filepath.Join(ToolDir, configFilename)
	if Verbose {
		log.Println("Reading config from ", path)
	}
	if _, err := os.Stat(ToolDir); os.IsNotExist(err) {
		if Verbose {
			log.Printf("Config dir not found, creating %s", ToolDir)
		}
		err := os.MkdirAll(ToolDir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if Verbose {
			log.Printf("Config file not found, creating default %s", path)
		}
		// TODO: Add default config & early return
		if err := os.WriteFile(path, []byte{}, 0644); err != nil {
			return err
		}
	}
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(raw, &Cfg)
	if err != nil {
		return err
	}
	return nil
}

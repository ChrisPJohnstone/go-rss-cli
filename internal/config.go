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

func configPath() string {
	return filepath.Join(ToolDir, configFilename)
}

func ReadConfig(cmd *cobra.Command, args []string) error {
	path := configPath()
	if Verbose {
		log.Println("Reading config from ", path)
	}
	ensureDir(ToolDir)
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

func WriteConfig(cmd *cobra.Command, args []string) error {
	path := configPath()
	if Verbose {
		log.Println("Writing config to ", path)
	}
	ensureDir(ToolDir)
	raw, err := yaml.Marshal(&Cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0644)
}

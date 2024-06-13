package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config struct
type Config struct {
	ApiEndpoint string `yaml:"api_endpoint"`
	ApiKey      string `yaml:"api_key"`
	ProfileID   string `yaml:"profile_id"`
}

// getConfigFile returns the path to the config file
func getConfigFile() string {
	return fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".config/bank/wise.yml")
}

// New returns a new configuration
func New() (Config, error) {
	var config Config

	// get config file
	configFile := getConfigFile()

	// open file
	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	// decode yaml
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return config, err
		}
	}

	return config, nil
}

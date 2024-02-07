package main

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-yaml/yaml"
)

// loadConfig loads configuration from a TOML file.
func loadConfig(path string) (Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func loadCities(path string) (Cities, error) {
	// Read the YAML file
	data, err := os.ReadFile("cities.yaml")
	if err != nil {
		return Cities{}, err
	}

	// Unmarshal the YAML into our struct
	var cities Cities
	err = yaml.Unmarshal(data, &cities)
	if err != nil {
		return Cities{}, err
	}

	return cities, nil
}

package main

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type OpenApiConfiguration struct {
	ApiKey string `yaml:"key"`
}

func LoadConfiguration(envFile string) (*OpenApiConfiguration, error) {
	filename, _ := filepath.Abs(envFile)
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &OpenApiConfiguration{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

package kopenai

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Openai OpenAiConfig       `yaml:"openai"`
	Naver  NaverOpenApiConfig `yaml:"naver"`
}

type OpenAiConfig struct {
	ApiKey string `yaml:"apiKey"`
}

type NaverOpenApiConfig struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
}

func LoadConfiguration(envFile string) (*Config, error) {
	filename, _ := filepath.Abs(envFile)
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

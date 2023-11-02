package kopenaigpt

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type AiConfig struct {
	Openai struct {
		ApiKey string `yaml:"apiKey"`
	} `yaml:"openai"`
	Naver NaverOpenApiConfig `yaml:"naver"`
}

type NaverOpenApiConfig struct {
	ClientId     string `yaml:"clientId"`
	ClientSecret string `yaml:"clientSecret"`
}

func LoadConfiguration(envFile string) (*AiConfig, error) {
	filename, _ := filepath.Abs(envFile)
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	config := &AiConfig{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

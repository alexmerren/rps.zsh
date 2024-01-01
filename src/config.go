package src

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	githubTokenKey  = "GITHUB_TOKEN"
	defaultFilename = "usr/local/share/rps/config.yaml"
)

type Config struct {
	Token string
}

func NewConfig() *Config {
	configData, err := os.ReadFile(defaultFilename)
	if err != nil {
		return nil
	}

	configMap := map[string]string{}
	if err = yaml.Unmarshal(configData, &configMap); err != nil {
		return nil
	}

	return &Config{
		Token: configMap[githubTokenKey],
	}
}

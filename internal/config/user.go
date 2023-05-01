package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	defaultConfigDirectory = "/usr/local/share/rps/"
	defaultConfigName      = "config.yaml"
)

var defaultConfigFilepath = fmt.Sprintf("%s%s", defaultConfigDirectory, defaultConfigName)

func generateUserHomeConfigPath() (string, error) {
	return fmt.Sprintf("%s%s", defaultConfigDirectory, defaultConfigName), nil
}

func CreateUserConfig() (*GithubConfig, error) {
	configFilepath, err := generateUserHomeConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(configFilepath); err != nil {
		return nil, fmt.Errorf("could not find config.yaml. Is it located in %s?", defaultConfigFilepath)
	}

	githubConfig := NewGithubConfig(configFilepath)
	if githubConfig == nil {
		return nil, errors.New("the config could not be read. Is it properly formatted?")
	}

	return githubConfig, nil
}

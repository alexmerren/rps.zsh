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

var (
	errConfigLocation = errors.New("could not find config")
	errConfigFormat   = errors.New("could not read config")
	//nolint:gochecknoglobals // I found an error, it always highlights this as a global variable when it is **not**.
	defaultConfigFilepath = fmt.Sprintf("%s%s", defaultConfigDirectory, defaultConfigName)
)

func generateUserHomeConfigPath() string {
	return defaultConfigFilepath
}

func CreateUserConfig() (*GithubConfig, error) {
	configFilepath := generateUserHomeConfigPath()

	if _, err := os.Stat(configFilepath); err != nil {
		return nil, fmt.Errorf("%w: %s", errConfigLocation, defaultConfigFilepath)
	}

	githubConfig := NewGithubConfig(configFilepath)
	if githubConfig == nil {
		return nil, errConfigFormat
	}

	return githubConfig, nil
}

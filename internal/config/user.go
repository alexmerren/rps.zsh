package config

import (
	"errors"
	"fmt"
	"os"
	"os/user"
)

const (
	defaultConfigName      = "config.yaml"
	defaultConfigDirectory = "/.config/rps/"
)

func generateUserHomeConfigPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	homeDirectory := currentUser.HomeDir
	configPath := fmt.Sprintf("%s%s%s", homeDirectory, defaultConfigDirectory, defaultConfigName)
	return configPath, nil
}

func CreateUserConfig() (*GithubConfig, error) {
	configFilepath, err := generateUserHomeConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err = os.Stat(configFilepath); err != nil {
		return nil, errors.New("could not find config.yaml. Is it located in $HOME/.config/rps?")
	}
	githubConfig := NewGithubConfig(configFilepath)
	if githubConfig == nil {
		return nil, errors.New("the config could not be read. Is it properly formatted?")
	}
	return githubConfig, nil
}

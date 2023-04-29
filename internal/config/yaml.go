package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	githubUsernameKey = "GITHUB_USERNAME"
	githubTokenKey    = "GITHUB_TOKEN"
)

type GithubConfig struct {
	token    string
	username string
}

func NewGithubConfig(filename string) *GithubConfig {
	configData, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	configMap := map[string]string{}
	if err = yaml.Unmarshal(configData, &configMap); err != nil {
		return nil
	}

	return &GithubConfig{
		token:    configMap[githubTokenKey],
		username: configMap[githubUsernameKey],
	}
}

func (g *GithubConfig) GetToken() string {
	return g.token
}

func (g *GithubConfig) GetUsername() string {
	return g.username
}

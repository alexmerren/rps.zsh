package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const (
	//nolint:gosec //These are not credentials.
	githubTokenKey = "GITHUB_TOKEN"
)

type GithubConfig struct {
	token string
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
		token: configMap[githubTokenKey],
	}
}

func (g *GithubConfig) GetToken() string {
	return g.token
}

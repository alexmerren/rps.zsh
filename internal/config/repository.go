package config

type GithubConfigurator interface {
	GetToken() (string, error)
}

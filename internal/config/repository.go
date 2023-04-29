package config

type GithubConfigurator interface {
	GetUsername() (string, error)
	GetToken() (string, error)
}

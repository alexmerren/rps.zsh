package client

type GithubInteractor interface {
	GetUserRepositories() ([]byte, error)
	GetStarredRepositories() ([]byte, error)
}

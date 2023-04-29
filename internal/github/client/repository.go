package client

type GithubInteractor interface {
	GetUserRepositories(username string) ([]byte, error)
	GetStarredRepositories(username string) ([]byte, error)
}

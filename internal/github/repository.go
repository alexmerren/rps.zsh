package github

import (
	"github.com/alexmerren/rps/internal/github/repository"
)

type API interface {
	GetUserRepositories(username string) ([]*repository.Repository, error)
	GetStarredRepositories(username string) ([]*repository.Repository, error)
}

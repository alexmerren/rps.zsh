package github

import (
	"context"

	"github.com/alexmerren/rps/internal/github/repository"
)

type GithubApi interface {
	GetUserRepositories(username string) ([]*repository.Repository, error)
	GetUserRepositoriesWithContext(ctx context.Context, username string) ([]*repository.Repository, error)
	GetStarredRepositories(username string) ([]*repository.Repository, error)
	GetStarredRepositoriesWithContext(ctx context.Context, username string) ([]*repository.Repository, error)
}

package github

import (
    "context"
)

type GithubApi interface { 
    GetUserRepositories(username string) ([]*Repository, error)
    GetUserRepositoriesWithContext(ctx context.Context, username string) ([]*Repository, error)
    GetStarredRepositories(username string) ([]*Repository, error)
    GetStarredRepositoriesWithContext(ctx context.Context, username string) ([]*Repository, error)
}

type GithubClient interface  {
    GetUserRepositories(username string) ([]byte, error)
    GetStarredRepositories(username string) ([]byte, error)
}

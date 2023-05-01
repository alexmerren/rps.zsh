package prompt

import "github.com/alexmerren/rps/internal/github/repository"

type RepositoryPrompter interface {
	SelectRepositoryPrompt(repositories []*repository.Repository) (int, error)
}

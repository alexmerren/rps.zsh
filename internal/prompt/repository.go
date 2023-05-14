package prompt

import (
	"io"

	"github.com/alexmerren/rps/internal/github/repository"
)

type RepositoryPrompter interface {
	SelectRepositoryPrompt(
		repositories []*repository.Repository,
		isVimMode bool,
		numLinesInPrompt int,
		stdout io.WriteCloser,
	) (int, error)
}

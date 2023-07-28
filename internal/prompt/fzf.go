package prompt

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexmerren/rps/internal/github/repository"
	fzf "github.com/ktr0731/go-fuzzyfinder"
)

func NewFzfPrompt(ctx context.Context, repositories []*repository.Repository) (int, error) {
	index, err := fzf.Find(repositories, func(index int) string {
		return fmt.Sprintf("%s/%s", repositories[index].GetOwner(), repositories[index].GetName())
	})

	if errors.Is(err, fzf.ErrAbort) {
		return 0, ErrPromptInterrupted
	}

	return index, err
}

package prompt

import (
	"context"
	"fmt"

	"github.com/alexmerren/rps/internal/github/repository"
	fzf "github.com/ktr0731/go-fuzzyfinder"
)

func NewFzfPrompt(ctx context.Context, repositories []*repository.Repository) (int, error) {
	return fzf.Find(repositories, func(index int) string {
		return fmt.Sprintf("%s/%s", repositories[index].GetOwner(), repositories[index].GetName())
	})
}

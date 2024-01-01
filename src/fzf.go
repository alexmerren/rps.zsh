package src

import (
	"errors"
	"fmt"

	fzf "github.com/ktr0731/go-fuzzyfinder"
)

var (
	ErrPromptInterrupt = errors.New("prompt has been interrupted")
)

func NewFzfPrompt(repositories []*Repository) (int, error) {
	index, err := fzf.Find(repositories, func(index int) string {
		return fmt.Sprintf("%s/%s", repositories[index].User, repositories[index].Name)
	})

	if errors.Is(err, fzf.ErrAbort) {
		return 0, ErrPromptInterrupt
	}

	return index, err
}

package prompt

import (
	"fmt"
	"io"
	"strings"

	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/manifoldco/promptui"
)

const defaultHideHelp = true

type GithubRepositoryPrompt struct{}

func NewGithubRepositoryPrompt() *GithubRepositoryPrompt {
	return &GithubRepositoryPrompt{}
}

func (g *GithubRepositoryPrompt) SelectRepositoryPrompt(
	repositories []*repository.Repository,
	isVimMode bool,
	numLinesInPrompt int,
	stdout io.WriteCloser,
) (int, error) {
	//nolint:exhaustruct // These do not all need to be declared.
	prompt := promptui.Select{
		Items:     repositories,
		Stdout:    stdout,
		HideHelp:  defaultHideHelp,
		IsVimMode: isVimMode,
		Size:      numLinesInPrompt,
		Templates: generateRepositoryTemplates(),
		Searcher:  createSearchingFunction(repositories),
	}
	index, _, err := prompt.Run()

	return index, fmt.Errorf("error in select prompt: %w", err)
}

func createSearchingFunction(repositories []*repository.Repository) func(string, int) bool {
	return func(input string, index int) bool {
		repository := repositories[index]
		name := strings.ReplaceAll(strings.ToLower(repository.GetName()), " ", "")
		owner := strings.ReplaceAll(strings.ToLower(repository.GetOwner()), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(name, input) || strings.Contains(owner, input)
	}
}

func generateRepositoryTemplates() *promptui.SelectTemplates {
	//nolint:exhaustruct // These do not all need to be declared.
	return &promptui.SelectTemplates{
		Label:    "Select a repository to download:",
		Active:   "\U00002705\t{{ .GetName | bold | green }} ({{ .GetOwner | bold | green }})",
		Inactive: " \t{{ .GetName | red }} ({{ .GetOwner | red }})",
		Selected: "\U00002705\t{{ .GetName | bold | green }}\U0000002F{{ .GetOwner | green }}",
	}
}

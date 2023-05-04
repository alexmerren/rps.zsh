package prompt

import (
	"io"
	"strings"

	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/manifoldco/promptui"
)

type GithubRepositoryPrompt struct{}

func NewGithubRepositoryPrompt() *GithubRepositoryPrompt {
	return &GithubRepositoryPrompt{}
}

func (g *GithubRepositoryPrompt) SelectRepositoryPrompt(repositories []*repository.Repository, isVimMode bool, numLinesInPrompt int, stdout io.WriteCloser) (int, error) {
	prompt := promptui.Select{
		Label:     "repository",
		Items:     repositories,
		Stdout:    stdout,
		IsVimMode: isVimMode,
		Size:      numLinesInPrompt,
		Templates: generateRepositoryTemplates(),
		Searcher:  createSearchingFunction(repositories),
	}
	index, _, err := prompt.Run()
	if err != nil {
		return 0, err
	}
	return index, err
}

func createSearchingFunction(repositories []*repository.Repository) func(string, int) bool {
	return func(input string, index int) bool {
		repository := repositories[index]
		name := strings.Replace(strings.ToLower(repository.GetName()), " ", "", -1)
		owner := strings.Replace(strings.ToLower(repository.GetOwner()), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(owner, input)
	}
}

func generateRepositoryTemplates() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "Choose a {{ . }} to download:",
		Active:   "\U00002705\t{{ .GetName | bold | green }} ({{ .GetOwner | bold | green }})",
		Inactive: " \t{{ .GetName | red }} ({{ .GetOwner | red }})",
		Selected: "\U00002705\t{{ .GetName | bold | green }}\U0000002F{{ .GetOwner | green }}",
	}
}

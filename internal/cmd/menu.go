package cmd

import (
	"context"
	"os/exec"
	"strings"

	"github.com/alexmerren/rps/internal/config"
	"github.com/alexmerren/rps/internal/github"
	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	defaultProtocol = "ssh"
)

func NewCmdMenu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "menu",
		Short: "Select repositories to manage",
		RunE: func(cmd *cobra.Command, args []string) error {
			repositories, err := getRepositories()
			if err != nil {
				return err
			}

			prompt, err := createPrompt(repositories)
			if err != nil {
				return err
			}

			selectedIndex, _, err := prompt.Run()
			if err != nil {
				return err
			}

			remoteUrl := repository.GenerateRepositoryRemoteUrl(repositories[selectedIndex], defaultProtocol)
			if err = downloadRepository(remoteUrl); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func downloadRepository(remoteUrl string) error {
	_, err := exec.Command("git", "clone", remoteUrl).Output()
	if err != nil {
		return err
	}
	return nil
}

func getRepositories() ([]*repository.Repository, error) {
	githubConfig := config.NewGithubConfig("./_config.yaml")
	token := githubConfig.GetToken()
	username := githubConfig.GetUsername()
	ctx := context.Background()
	client := client.NewGithubClientWithAuthentication(token)
	api := github.NewGithubUserApi(ctx, client)
	starredRepositories, err := api.GetStarredRepositories(username)
	if err != nil {
		return nil, err
	}

	userRepositories, err := api.GetUserRepositories(username)
	if err != nil {
		return nil, err
	}

	repositories := append(userRepositories, starredRepositories...)
	return repositories, nil
}

func createPrompt(repositories []*repository.Repository) (promptui.Select, error) {
	searchingFunction := func(input string, index int) bool {
		repository := repositories[index]
		name := strings.Replace(strings.ToLower(repository.GetName()), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		owner := strings.Replace(strings.ToLower(repository.GetOwner()), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(owner, input)
	}

	templates := &promptui.SelectTemplates{
		Label:    "Choose a {{ . }} to download:",
		Active:   "\U00002705\t{{ .GetName | bold | green }} ({{ .GetOwner | bold | green }})",
		Inactive: " \t{{ .GetName | red }} ({{ .GetOwner | red }})",
		Selected: "\U00002705\t{{ .GetName | bold | green }}\U0000002F{{ .GetOwner | green }}",
	}

	prompt := promptui.Select{
		Label:     "repository",
		Items:     repositories,
		Templates: templates,
		Size:      25,
		Searcher:  searchingFunction,
	}
	return prompt, nil
}

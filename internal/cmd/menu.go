package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/alexmerren/rps/internal/config"
	"github.com/alexmerren/rps/internal/github"
	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	defaultProtocol     = "ssh"
	configFileName      = "config.yaml"
	configFileDirectory = "/.config/rps/"
)

func NewCmdMenu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "menu",
		Short: "Select repositories to manage",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			configPath, err := generateConfigPath()
			if err != nil {
				return err
			}

			if _, err = os.Stat(configPath); err != nil {
				return errors.New("could not find config.yaml. Is it located in $HOME/.config/rps?")
			}

			repositories, err := getRepositories(ctx, configPath)
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
			if err = downloadRepository(ctx, remoteUrl); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func generateConfigPath() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	homeDirectory := currentUser.HomeDir
	configPath := fmt.Sprintf("%s%s%s", homeDirectory, configFileDirectory, configFileName)
	return configPath, nil
}

func getRepositories(ctx context.Context, configPath string) ([]*repository.Repository, error) {
	githubConfig := config.NewGithubConfig(configPath)
	if githubConfig == nil {
		return nil, errors.New("the config could not be read. Is it properly formatted?")
	}
	token := githubConfig.GetToken()
	username := githubConfig.GetUsername()
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

func downloadRepository(ctx context.Context, remoteUrl string) error {
	out, err := exec.CommandContext(ctx, "git", "clone", remoteUrl).Output()
	if err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, string(out[:]))
	return nil
}

package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/alexmerren/rps/internal/config"
	"github.com/alexmerren/rps/internal/github"
	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/alexmerren/rps/internal/prompt"
	"github.com/spf13/cobra"
)

var errConfigNotValid = errors.New("could not find token or username")

const (
	defaultProtocol  = false
	protocolNameLong = "http"
	protocolUsage    = "clone over https instead of ssh"
)

func NewCmdRoot() *cobra.Command {
	//nolint:exhaustruct // All of this does not need to be set.
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Select repositories to download",
		RunE: func(cmd *cobra.Command, args []string) error {
			protocol, err := cmd.Flags().GetBool(protocolNameLong)
			if err != nil {
				return fmt.Errorf("could not get protocol flag: %w", err)
			}

			protocolString := ""
			if protocol {
				protocolString = "https"
			} else {
				protocolString = "ssh"
			}

			return rootRun(cmd.Context(), protocolString)
		},
	}
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.PersistentFlags().Bool(protocolNameLong, defaultProtocol, protocolUsage)

	return cmd
}

func rootRun(ctx context.Context, remoteProtocol string) error {
	config, err := config.CreateUserConfig()
	if err != nil {
		return fmt.Errorf("could not get user config: %w", err)
	}

	repositories, err := getRepositoriesWithConfig(config)
	if err != nil {
		return err
	}

	selectedIndex, err := prompt.NewFzfPrompt(repositories)
	if err != nil {
		return fmt.Errorf("error in prompt: %w", err)
	}

	remoteURL := repository.GenerateRepositoryRemoteURL(repositories[selectedIndex], remoteProtocol)
	err = github.CallOsGitClone(ctx, remoteURL)

	return err
}

func getRepositoriesWithConfig(config *config.GithubConfig) ([]*repository.Repository, error) {
	token := config.GetToken()

	if token == "" {
		return nil, errConfigNotValid
	}

	client := client.NewGithubClientWithAuthentication(token)
	api := github.NewUserAPI(client)

	starredRepositories, err := api.GetStarredRepositories()
	if err != nil {
		return nil, fmt.Errorf("could not get starred repositories: %w", err)
	}

	userRepositories, err := api.GetUserRepositories()
	if err != nil {
		return nil, fmt.Errorf("could not get user repositories: %w", err)
	}

	userRepositories = append(userRepositories, starredRepositories...)

	return userRepositories, nil
}

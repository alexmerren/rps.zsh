package cmd

import (
	"context"
	"errors"

	"github.com/alexmerren/rps/internal/config"
	"github.com/alexmerren/rps/internal/github"
	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/alexmerren/rps/internal/prompt"
	"github.com/spf13/cobra"
)

const (
	defaultProtocol = "ssh"

	defaultIsVimMode     = false
	vimModeFlagNameLong  = "vimmode"
	vimModeFlagNameShort = "v"
	vimModeFlagUsage     = ""
)

func NewCmdMenu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Select repositories to download",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := config.CreateUserConfig()
			if err != nil {
				return err
			}

			repositories, err := getRepositoriesWithConfig(cmd.Context(), config)
			if err != nil {
				return err
			}

			isVimMode, err := cmd.Flags().GetBool(vimModeFlagNameLong)
			if err != nil {
				return err
			}

			prompter := prompt.NewGithubRepositoryPrompt()
			selectedIndex, err := prompter.SelectRepositoryPrompt(repositories, isVimMode)
			if err != nil {
				return err
			}

			remoteUrl := repository.GenerateRepositoryRemoteUrl(repositories[selectedIndex], defaultProtocol)
			if err = github.CallOsGitClone(cmd.Context(), remoteUrl); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolP(vimModeFlagNameLong, vimModeFlagNameShort, defaultIsVimMode, vimModeFlagUsage)
	return cmd
}

func getRepositoriesWithConfig(ctx context.Context, config *config.GithubConfig) ([]*repository.Repository, error) {
	token := config.GetToken()
	username := config.GetUsername()
	if token == "" || username == "" {
		return nil, errors.New("could not find token or username")
	}
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

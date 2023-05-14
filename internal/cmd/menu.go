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
	defaultProtocol = "ssh"

	defaultIsVimMode     = false
	vimModeFlagNameLong  = "vimmode"
	vimModeFlagNameShort = "v"
	vimModeFlagUsage     = "use vim-like controls on the selection prompt"

	defaultNumLines       = 10
	numLinesFlagNameLong  = "lines"
	numLinesFlagNameShort = "l"
	numLinesFlagUsage     = "number of lines to display on the selection prompt"
)

func NewCmdRoot() *cobra.Command {
	//nolint:exhaustruct // All of this does not need to be set.
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Select repositories to download",
		RunE: func(cmd *cobra.Command, args []string) error {
			isVimMode, err := cmd.Flags().GetBool(vimModeFlagNameLong)
			if err != nil {
				return fmt.Errorf("could not get vim mode flag: %w", err)
			}

			numberOfLines, err := cmd.Flags().GetInt(numLinesFlagNameLong)
			if err != nil {
				return fmt.Errorf("could not get number of lines flag: %w", err)
			}

			return rootRun(cmd.Context(), isVimMode, numberOfLines, defaultProtocol)
		},
	}
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.PersistentFlags().BoolP(vimModeFlagNameLong, vimModeFlagNameShort, defaultIsVimMode, vimModeFlagUsage)
	cmd.PersistentFlags().IntP(numLinesFlagNameLong, numLinesFlagNameShort, defaultNumLines, numLinesFlagUsage)

	return cmd
}

func rootRun(ctx context.Context, isVimMode bool, numLinesInPrompt int, remoteProtocol string) error {
	config, err := config.CreateUserConfig()
	if err != nil {
		return fmt.Errorf("could not get user config: %w", err)
	}

	repositories, err := getRepositoriesWithConfig(config)
	if err != nil {
		return err
	}

	prompter := prompt.NewGithubRepositoryPrompt()
	bellSkipperStdout := prompt.NewBellSkipperStdout()

	selectedIndex, err := prompter.SelectRepositoryPrompt(repositories, isVimMode, numLinesInPrompt, bellSkipperStdout)
	if err != nil {
		return fmt.Errorf("error in prompt: %w", err)
	}

	remoteURL := repository.GenerateRepositoryRemoteURL(repositories[selectedIndex], remoteProtocol)
	err = github.CallOsGitClone(ctx, remoteURL)

	//nolint:wrapcheck // This does not need to be wrapped.
	return err
}

func getRepositoriesWithConfig(config *config.GithubConfig) ([]*repository.Repository, error) {
	token := config.GetToken()
	username := config.GetUsername()

	if token == "" || username == "" {
		return nil, errConfigNotValid
	}

	client := client.NewGithubClientWithAuthentication(token)
	api := github.NewUserAPI(client)

	starredRepositories, err := api.GetStarredRepositories(username)
	if err != nil {
		return nil, fmt.Errorf("could not get starred repositories: %w", err)
	}

	userRepositories, err := api.GetUserRepositories(username)
	if err != nil {
		return nil, fmt.Errorf("could not get user repositories: %w", err)
	}

	userRepositories = append(userRepositories, starredRepositories...)

	return userRepositories, nil
}

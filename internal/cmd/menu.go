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
	vimModeFlagUsage     = "use vim-like controls on the selection prompt"

	defaultNumLines       = 10
	numLinesFlagNameLong  = "lines"
	numLinesFlagNameShort = "l"
	numLinesFlagUsage     = "number of lines to display on the selection prompt"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Select repositories to download",
		RunE: func(cmd *cobra.Command, args []string) error {
			isVimMode, err := cmd.Flags().GetBool(vimModeFlagNameLong)
			if err != nil {
				return err
			}

			numberOfLines, err := cmd.Flags().GetInt(numLinesFlagNameLong)
			if err != nil {
				return err
			}

			return rootRun(cmd.Context(), isVimMode, numberOfLines, defaultProtocol)
		},
	}
	cmd.SilenceErrors = true // This removes the command output from terminal
	cmd.SilenceUsage = true  // this removes the usage output when cancelling
	cmd.PersistentFlags().BoolP(vimModeFlagNameLong, vimModeFlagNameShort, defaultIsVimMode, vimModeFlagUsage)
	cmd.PersistentFlags().IntP(numLinesFlagNameLong, numLinesFlagNameShort, defaultNumLines, numLinesFlagUsage)

	return cmd
}

func rootRun(ctx context.Context, isVimMode bool, numLinesInPrompt int, remoteProtocol string) error {
	config, err := config.CreateUserConfig()
	if err != nil {
		return err
	}

	repositories, err := getRepositoriesWithConfig(ctx, config)
	if err != nil {
		return err
	}

	prompter := prompt.NewGithubRepositoryPrompt()
	bellSkipperStdout := prompt.NewBellSkipperStdout()
	selectedIndex, err := prompter.SelectRepositoryPrompt(repositories, isVimMode, numLinesInPrompt, bellSkipperStdout)
	if err != nil {
		return err
	}

	remoteUrl := repository.GenerateRepositoryRemoteUrl(repositories[selectedIndex], remoteProtocol)
	if err = github.CallOsGitClone(ctx, remoteUrl); err != nil {
		return err
	}
	return nil
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

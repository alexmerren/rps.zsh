package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/alexmerren/rps/src"
	"github.com/spf13/cobra"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func main() {
	exitCode := mainRun()
	os.Exit(int(exitCode))
}

func mainRun() exitCode {
	ctx := context.Background()
	config := src.NewConfig()
	rpsCmd := NewRpsCommand(config)

	if _, err := rpsCmd.ExecuteContextC(ctx); err != nil && !errors.Is(err, src.ErrPromptInterrupt) {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		return exitError
	}

	return exitOK
}

func NewRpsCommand(config *src.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rps",
		Short: "Select a repository to clone",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			client := src.NewClient(config.Token)
			repositories, err := src.GetUserRepositories(ctx, client)
			if err != nil {
				return err
			}

			selectedIndex, err := src.NewFzfPrompt(repositories)
			if err != nil {
				return fmt.Errorf("error in prompt: %w", err)
			}

			return repositories[selectedIndex].Clone(ctx)
		},
	}
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true

	return cmd
}

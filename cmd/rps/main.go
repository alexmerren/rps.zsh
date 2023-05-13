package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/alexmerren/rps/internal/cmd"
	"github.com/alexmerren/rps/internal/prompt"
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
	rootCmd := cmd.NewCmdRoot()

	if _, err := rootCmd.ExecuteContextC(ctx); err != nil && !errors.Is(err, prompt.ErrPromptInterrupted) {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())

		return exitError
	}

	return exitOK
}

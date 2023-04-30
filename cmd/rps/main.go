package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alexmerren/rps/internal/cmd"
)

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

const version = "0.1.0"

func main() {
	exitCode := mainRun()
	os.Exit(int(exitCode))
}

func mainRun() exitCode {
	ctx := context.Background()
	menuCmd := cmd.NewCmdMenu()
	if _, err := menuCmd.ExecuteContextC(ctx); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return exitError
	}
	return exitOK
}

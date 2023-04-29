package main

import (
	"context"
	"os"
    "fmt"

	"github.com/alexmerren/repo-selector/internal/config"
	"github.com/alexmerren/repo-selector/internal/github"
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
    githubConfig := config.NewGithubConfig("./_config.yaml")
    token := githubConfig.GetToken()
    username := githubConfig.GetUsername()
    ctx := context.Background()
    client := github.NewGithubAuthClient(token)
    api := github.NewGithubUserApi(ctx, client)
    starredRepositories, err := api.GetStarredRepositories(username)
    if err != nil {
        return exitError
    }
    userRepositories, err := api.GetUserRepositories(username)
    if err != nil {
        return exitError
    }
    allRepositories := append(userRepositories, starredRepositories...)
    for _, repository := range allRepositories {
        fmt.Println(repository)
    }
    return exitOK
}


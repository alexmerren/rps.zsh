package main

import (
	"os"
    "context"
    "fmt"

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
    ctx := context.Background()
    client := github.NewGithubNoAuthClient()
    api := github.NewGithubUserApi(ctx, client)
    repositories, err := api.GetStarredRepositories("alexmerren")
    if err != nil {
        fmt.Println(err)
        return exitError
    }
    for _, repo := range repositories {
        fmt.Println(repo)
    }
    return exitOK 
}


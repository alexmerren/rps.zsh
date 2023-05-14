package github

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/exec"

	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/buger/jsonparser"
)

type UserAPI struct {
	client client.GithubInteractor
}

func NewUserAPI(client client.GithubInteractor) *UserAPI {
	return &UserAPI{
		client: client,
	}
}

func (g *UserAPI) GetUserRepositories(username string) ([]*repository.Repository, error) {
	userRepositoriesRaw, err := g.client.GetUserRepositories(username)
	if err != nil {
		return nil, fmt.Errorf("could not get user repositories: %w", err)
	}

	userRepositories, err := GetRepositoriesFromRaw(userRepositoriesRaw)
	if err != nil {
		return nil, fmt.Errorf("could not format repositories: %w", err)
	}

	return userRepositories, nil
}

func (g *UserAPI) GetStarredRepositories(username string) ([]*repository.Repository, error) {
	starredRepositoriesRaw, err := g.client.GetStarredRepositories(username)
	if err != nil {
		return nil, fmt.Errorf("could not get starred repositories: %w", err)
	}

	starredRepositories, err := GetRepositoriesFromRaw(starredRepositoriesRaw)
	if err != nil {
		return nil, fmt.Errorf("could not format repositories: %w", err)
	}

	return starredRepositories, nil
}

func GetRepositoriesFromRaw(raw []byte) ([]*repository.Repository, error) {
	repositories := make([]*repository.Repository, 0)

	_, err := jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rawURL, _ := jsonparser.GetString(value, "html_url")
		parsedURL, _ := url.Parse(rawURL)
		newRepository, _ := repository.NewRepositoryFromURL(parsedURL)
		repositories = append(repositories, newRepository)
	})
	if err != nil {
		return nil, fmt.Errorf("error in parsing JSON: %w", err)
	}

	return repositories, nil
}

func CallOsGitClone(ctx context.Context, remoteURL string) error {
	cmd := exec.CommandContext(ctx, "git", "clone", remoteURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	return fmt.Errorf("error calling git clone: %w", err)
}

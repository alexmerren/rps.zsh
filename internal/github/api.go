package github

import (
	"context"
	"strings"

	"github.com/alexmerren/rps/internal/github/client"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/buger/jsonparser"
)

type GithubUserApi struct {
	ctx    context.Context
	client client.GithubInteractor
}

func NewGithubUserApi(ctx context.Context, client client.GithubInteractor) *GithubUserApi {
	return &GithubUserApi{
		ctx:    ctx,
		client: client,
	}
}

func (g *GithubUserApi) GetUserRepositories(username string) ([]*repository.Repository, error) {
	return g.GetUserRepositoriesWithContext(g.ctx, username)
}

func (g *GithubUserApi) GetUserRepositoriesWithContext(ctx context.Context, username string) ([]*repository.Repository, error) {
	userRepositoriesRaw, err := g.client.GetUserRepositories(username)
	if err != nil {
		return nil, err
	}
	userRepositories, err := getRepositoriesFromRaw(userRepositoriesRaw)
	if err != nil {
		return nil, err
	}
	return userRepositories, nil
}

func (g *GithubUserApi) GetStarredRepositories(username string) ([]*repository.Repository, error) {
	return g.GetStarredRepositoriesWithContext(g.ctx, username)
}

func (g *GithubUserApi) GetStarredRepositoriesWithContext(ctx context.Context, username string) ([]*repository.Repository, error) {
	starredRepositoriesRaw, err := g.client.GetStarredRepositories(username)
	if err != nil {
		return nil, err
	}
	starredRepositories, err := getRepositoriesFromRaw(starredRepositoriesRaw)
	if err != nil {
		return nil, err
	}
	return starredRepositories, nil
}

func getRepositoriesFromRaw(raw []byte) ([]*repository.Repository, error) {
	repositories := make([]*repository.Repository, 0)
	_, err := jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		name, _ := jsonparser.GetString(value, "name")
		owner, _ := jsonparser.GetString(value, "owner", "login")
		rawHost, _ := jsonparser.GetString(value, "html_url")
		host := strings.SplitN(rawHost[8:], "/", 2)[0]
		repositories = append(repositories, repository.NewRepositoryWithHost(owner, name, host))
	})
	if err != nil {
		return nil, err
	}
	return repositories, nil
}

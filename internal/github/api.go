package github

import (
	"context"
	"net/url"

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
	userRepositories, err := GetRepositoriesFromRaw(userRepositoriesRaw)
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
	starredRepositories, err := GetRepositoriesFromRaw(starredRepositoriesRaw)
	if err != nil {
		return nil, err
	}
	return starredRepositories, nil
}

func GetRepositoriesFromRaw(raw []byte) ([]*repository.Repository, error) {
	repositories := make([]*repository.Repository, 0)
	_, err := jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rawUrl, _ := jsonparser.GetString(value, "html_url")
		parsedUrl, _ := url.Parse(rawUrl)
		newRepository, _ := repository.NewRepositoryFromUrl(parsedUrl)
		repositories = append(repositories, newRepository)
	})
	if err != nil {
		return nil, err
	}
	return repositories, nil
}

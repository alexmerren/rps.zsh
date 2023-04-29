package github

import (
	"context"
    "fmt"
    "github.com/buger/jsonparser"
)

type GithubUserApi struct {
    ctx context.Context
    client GithubClient
}

func NewGithubUserApi(ctx context.Context, client GithubClient) *GithubUserApi {
    return &GithubUserApi{
        ctx: ctx,
        client: client,
    }
}

func (g *GithubUserApi) GetUserRepositories(username string) ([]*Repository, error) {
    return g.GetUserRepositoriesWithContext(g.ctx, username)
}

func (g *GithubUserApi) GetUserRepositoriesWithContext(ctx context.Context, username string) ([]*Repository, error) {
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

func (g *GithubUserApi) GetStarredRepositories(username string) ([]*Repository, error) {
    return g.GetStarredRepositoriesWithContext(g.ctx, username)
}

func (g *GithubUserApi) GetStarredRepositoriesWithContext(ctx context.Context, username string) ([]*Repository, error) {
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

func getRepositoriesFromRaw(raw []byte) ([]*Repository, error) {
    repositories := make([]*Repository, 0)
    _, err := jsonparser.ArrayEach(raw,  func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
        name, _:= jsonparser.GetString(value, "name")
        fmt.Println(name)
        owner, _:= jsonparser.GetString(value, "owner", "login")
        hostname, _:= jsonparser.GetString(value, "html_url")
        repositories= append(repositories, NewRepositoryWithHost(owner, name, hostname[8:18]))
    })
    if err != nil {
        return nil, err
    }
    return repositories, nil
}

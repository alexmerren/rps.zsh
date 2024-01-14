package src

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	maxPerPage       = 100
	visibilityAll    = "all"
	visibilityPublic = "public"
)

type Client struct {
	ctx          context.Context
	githubClient *github.Client
}

func NewClient(ctx context.Context, token string) *Client {
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: token,
		},
	)
	authClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := github.NewClient(authClient)

	return &Client{
		ctx:          ctx,
		githubClient: githubClient,
	}
}

func (c *Client) ListUserRepos() {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: maxPerPage},
		Visibility:  visibilityAll,
	}

	for {
		repos, resp, err := c.githubClient.Repositories.List(c.ctx, "", opt)
		if err != nil {
			fmt.Println(err)
			return
		}

		printRepos(repos)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}

func (c *Client) ListOrgRepos(organization string) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: maxPerPage},
		Type:        visibilityPublic,
	}

	for {
		repos, resp, err := c.githubClient.Repositories.ListByOrg(c.ctx, organization, opt)
		if err != nil {
			fmt.Println(err)
			return
		}

		printRepos(repos)

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}

func printRepos(repos []*github.Repository) {
	for _, repo := range repos {
		fmt.Println(*repo.FullName)
	}
}

package github

import (
    "net/url"
    "fmt"
    "strings"
)

const (
    githubHost = "github.com"
)

type Repository struct {
    Owner string
    Name string
    Host string
}

func NewRepository(owner, name string) *Repository {
    return &Repository{
        Owner: owner,
        Name: name,
        Host: githubHost,
    }
}

func NewRepositoryWithHost(owner, name, hostname string) *Repository {
    return &Repository{
        Owner: owner,
        Name: name,
        Host: hostname,
    }
}

func NewRepositoryFromUrl(url *url.URL) (*Repository, error) {
	if url.Hostname() == "" {
		return nil, fmt.Errorf("no hostname detected")
	}

	parts := strings.SplitN(strings.Trim(url.Path, "/"), "/", 3)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid path: %s", url.Path)
	}

	return NewRepositoryWithHost(parts[0], strings.TrimSuffix(parts[1], ".git"), url.Hostname()), nil
}

func GenerateRepositoryUrl(repository *Repository) (string, error) {
    return "", nil
}

func GenerateRepositoryRemoteUrl(repository *Repository, protocol string) (string, error) {
    if protocol == "ssh" {
		return fmt.Sprintf("git@%s:%s/%s.git", repository.Host, repository.Owner, repository.Name), nil
	}
    return fmt.Sprintf("https://%s/%s/%s", repository.Host, repository.Owner, repository.Name), nil
}

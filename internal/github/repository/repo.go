package repository

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

const (
	githubHost     = "github.com"
	numPartsInURL  = 2
	numSplitsInURL = 3
)

var (
	errNoHostnameInURL  = errors.New("no hostname detected")
	errInvalidPathInURL = errors.New("invalid path in URL")
)

type Repository struct {
	owner string
	name  string
	host  string
}

func NewRepository(owner, name string) *Repository {
	return &Repository{
		owner: owner,
		name:  name,
		host:  githubHost,
	}
}

func NewRepositoryWithHost(owner, name, host string) *Repository {
	return &Repository{
		owner: owner,
		name:  name,
		host:  host,
	}
}

func NewRepositoryFromURL(url *url.URL) (*Repository, error) {
	if url.Hostname() == "" {
		return nil, errNoHostnameInURL
	}

	parts := strings.SplitN(strings.Trim(url.Path, "/"), "/", numSplitsInURL)
	if len(parts) != numPartsInURL {
		return nil, fmt.Errorf("%w: %s", errInvalidPathInURL, url.Path)
	}

	return NewRepositoryWithHost(parts[0], strings.TrimSuffix(parts[1], ".git"), url.Hostname()), nil
}

func GenerateRepositoryRemoteURL(repository *Repository, protocol string) string {
	if protocol == "ssh" {
		return fmt.Sprintf("git@%s:%s/%s.git", repository.host, repository.owner, repository.name)
	}

	return fmt.Sprintf("https://%s/%s/%s", repository.host, repository.owner, repository.name)
}

func (r *Repository) GetOwner() string {
	return r.owner
}

func (r *Repository) GetName() string {
	return r.name
}

func (r *Repository) GetHost() string {
	return r.host
}

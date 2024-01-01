package src

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const (
	numPartsInURL  = 2
	numSplitsInURL = 3
)

var (
	errInvalidURL = errors.New("invalid repository URL")
)

type Repository struct {
	Host string
	User string
	Name string
}

func NewRepository(url *url.URL) (*Repository, error) {
	if url.Hostname() == "" {
		return nil, errInvalidURL
	}

	trimmedUrl := strings.Trim(url.Path, "/")
	urlParts := strings.SplitN(trimmedUrl, "/", numSplitsInURL)
	if len(urlParts) != numPartsInURL {
		return nil, fmt.Errorf("%w: %s", errInvalidURL, url.Path)
	}

	return &Repository{
		Host: urlParts[0],
		User: strings.TrimSuffix(urlParts[1], ".git"),
		Name: url.Hostname(),
	}, nil
}

func (r *Repository) Clone(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, "git", "clone", r.toRemoteUrl())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error calling git clone: %w", err)
	}

	return nil
}

func (r *Repository) toRemoteUrl() string {
	return fmt.Sprintf("git@%s:%s/%s.git", r.Host, r.User, r.Name)
}

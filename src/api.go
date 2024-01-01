package src

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"sync"

	"github.com/buger/jsonparser"
)

var (
	ErrEmptyRepositories = errors.New("no repositories found")
)

func GetUserRepositories(ctx context.Context, client *Client) ([]*Repository, error) {
	wg := &sync.WaitGroup{}
	repositories := make([]*Repository, 0)

	getRepositories(wg, repositories, client.GetUserRepositories)
	getRepositories(wg, repositories, client.GetStarredRepositories)

	wg.Wait()
	if len(repositories) == 0 {
		return nil, ErrEmptyRepositories
	}

	return repositories, nil
}

func getRepositories(wg *sync.WaitGroup, repositories []*Repository, getFunc func(int) ([]byte, error)) {
	moreRepositories := true
	pageNumber := 0

	for moreRepositories {
		wg.Add(1)
		pageNumber += 1
		go func() error {
			response, err := getFunc(pageNumber)
			if err != nil {
				return err
			}

			newRepositories, err := mapToRepositories(response)
			if err != nil {
				return err
			}

			repositories = append(repositories, newRepositories...)
			moreRepositories = hasNext(response)
			wg.Done()

			return nil
		}()
	}
}

func mapToRepositories(data []byte) ([]*Repository, error) {
	repositories := make([]*Repository, 0)

	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rawURL, _ := jsonparser.GetString(value, "html_url")
		parsedURL, _ := url.Parse(rawURL)
		newRepository, _ := NewRepository(parsedURL)
		repositories = append(repositories, newRepository)
	})
	if err != nil {
		return nil, fmt.Errorf("error in parsing JSON: %w", err)
	}

	return repositories, nil
}

// TODO Check if there are more pages to paginate
func hasNext(data []byte) bool {
	return false
}

// TODO Write this function
func GetOrgRepositories(ctx context.Context, client *Client, org string) ([]*Repository, error) {
	return nil, nil
}

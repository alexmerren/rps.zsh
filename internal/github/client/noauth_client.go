package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiURL                 = "https://api.github.com"
	listRepositoryEndpoint = "/users/%s/repos"
	listStarredEndpoint    = "/users/%s/starred"
	defaultAPIVersion      = "2022-11-28"
	versionHeaderName      = "X-GitHub-Api-Version"
)

var (
	errCreateListRequest     = errors.New("could not create list repository request")
	errCreateStarredRequest  = errors.New("could not create starred repository request")
	errIncorrectResponseCode = errors.New("incorrect response code")
)

type GithubClient struct {
	version string
	client  *http.Client
}

func NewGithubClient() *GithubClient {
	return &GithubClient{
		version: defaultAPIVersion,
		client:  http.DefaultClient,
	}
}

func (g *GithubClient) GetUserRepositories(username string) ([]byte, error) {
	request := createRequest(apiURL, listRepositoryEndpoint, username, g.version)
	if request == nil {
		return nil, errCreateListRequest
	}

	return doRequestAndReturnBody(request, g.client)
}

func (g *GithubClient) GetStarredRepositories(username string) ([]byte, error) {
	request := createRequest(apiURL, listStarredEndpoint, username, g.version)
	if request == nil {
		return nil, errCreateStarredRequest
	}

	return doRequestAndReturnBody(request, g.client)
}

func createRequest(apiURL, endpoint, username, version string) *http.Request {
	formattedEndpoint := fmt.Sprintf(endpoint, username)

	url, err := url.Parse(fmt.Sprintf("%s%s", apiURL, formattedEndpoint))
	if err != nil {
		return nil
	}

	headers := map[string][]string{
		versionHeaderName: {version},
	}

	//nolint:exhaustruct // All the properties do not need to be defined.
	return &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}
}

func doRequestAndReturnBody(request *http.Request, client *http.Client) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", errIncorrectResponseCode, response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	err = response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close response body: %w", err)
	}

	return responseBody, nil
}

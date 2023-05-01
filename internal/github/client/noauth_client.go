package client

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiUrl                 = "https://api.github.com"
	listRepositoryEndpoint = "/users/%s/repos"
	listStarredEndpoint    = "/users/%s/starred"
	defaultApiVersion      = "2022-11-28"
	versionHeaderName      = "X-GitHub-Api-Version"
)

type GithubClient struct {
	version string
	client  *http.Client
}

func NewGithubClient() *GithubClient {
	return &GithubClient{
		version: defaultApiVersion,
		client:  http.DefaultClient,
	}
}

func (g *GithubClient) GetUserRepositories(username string) ([]byte, error) {
	request := createRequest(apiUrl, listRepositoryEndpoint, username, g.version)
	if request == nil {
		return nil, errors.New("could not create authenticated list repository request")
	}
	return doRequestAndReturnBody(request, g.client)
}

func (g *GithubClient) GetStarredRepositories(username string) ([]byte, error) {
	request := createRequest(apiUrl, listStarredEndpoint, username, g.version)
	if request == nil {
		return nil, errors.New("could not create authenticated list starred request")
	}
	return doRequestAndReturnBody(request, g.client)
}

func createRequest(apiUrl, endpoint, username, version string) *http.Request {
	formattedEndpoint := fmt.Sprintf(endpoint, username)
	url, err := url.Parse(fmt.Sprintf("%s%s", apiUrl, formattedEndpoint))
	if err != nil {
		return nil
	}

	headers := map[string][]string{
		versionHeaderName: {version},
	}

	return &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}
}

func doRequestAndReturnBody(request *http.Request, client *http.Client) ([]byte, error) {
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %s", err.Error())
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println(response.StatusCode)
		return nil, fmt.Errorf("response code was %d", response.StatusCode)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %s", err.Error())
	}
	err = response.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to close response body: %s", err.Error())
	}
	return responseBody, nil
}

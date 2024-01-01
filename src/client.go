package src

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	apiUrl          = "https://api.github.com"
	userEndpointKey = "user/repos?sort=updated&direction=desc&per_page=100&page=%d"
	starEndpointKey = "user/starred?sort=updated&direction=desc&per_page=100&page=%d"
	orgEndpointKey  = "orgs/%s/repos?sort=updated&direction=desc&per_page=100&page=%d"

	defaultAPIVersion = "2022-11-28"
	authHeaderKey     = "Authorization"
	versionHeaderKey  = "X-GitHub-Api-Version"
)

var (
	ErrRequestAuth  = errors.New("could not create authenticated request")
	ErrResponseCode = errors.New("incorrect response code")
	ErrInvalidUrl   = errors.New("invalid request url")
)

type Client struct {
	token   string
	version string
	client  *http.Client
}

func NewClient(token string) *Client {
	if token == "" {
		return nil
	}

	return &Client{
		token:   token,
		version: defaultAPIVersion,
		client:  http.DefaultClient,
	}
}

func (c *Client) GetUserRepositories(pageNumber int) ([]byte, error) {
	endpoint := fmt.Sprintf(userEndpointKey, pageNumber)
	return c.do(endpoint)
}

func (c *Client) GetOrgRepositories(org string, pageNumber int) ([]byte, error) {
	endpoint := fmt.Sprintf(orgEndpointKey, org, pageNumber)
	return c.do(endpoint)
}

func (c *Client) GetStarredRepositories(pageNumber int) ([]byte, error) {
	endpoint := fmt.Sprintf(starEndpointKey, pageNumber)
	return c.do(endpoint)
}

func (c *Client) do(endpoint string) ([]byte, error) {
	url, err := url.Parse(fmt.Sprintf("%s/%s", apiUrl, endpoint))
	if err != nil {
		return nil, ErrInvalidUrl
	}

	tokenValue := fmt.Sprintf("Bearer %s", c.token)
	headers := map[string][]string{
		versionHeaderKey: {c.version},
		authHeaderKey:    {tokenValue},
	}

	request := &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrResponseCode, response.StatusCode)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	defer response.Body.Close()

	return responseBody, nil
}

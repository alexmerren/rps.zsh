package client

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	authListRepositoryEndpoint = "/user/repos"
	authListStarredEndpoint    = "/user/starred"
	authHeaderName             = "Authorization"
	authHeaderPrefix           = "Bearer %s"
)

var (
	errCreateAuthListRequest    = errors.New("could not create authenticated list repository request")
	errCreateAuthStarredRequest = errors.New("could not create authenticated starred repository request")
)

type GithubClientWithAuthentication struct {
	token   string
	version string
	client  *http.Client
}

func NewGithubClientWithAuthentication(token string) *GithubClientWithAuthentication {
	if token == "" {
		return nil
	}

	return &GithubClientWithAuthentication{
		token:   token,
		version: defaultAPIVersion,
		client:  http.DefaultClient,
	}
}

func (g *GithubClientWithAuthentication) GetUserRepositories(_ string) ([]byte, error) {
	request := createRequestWithAuthentication(apiURL, authListRepositoryEndpoint, g.token, g.version)
	if request == nil {
		return nil, errCreateAuthListRequest
	}

	return doRequestAndReturnBody(request, g.client)
}

func (g *GithubClientWithAuthentication) GetStarredRepositories(_ string) ([]byte, error) {
	request := createRequestWithAuthentication(apiURL, authListStarredEndpoint, g.token, g.version)
	if request == nil {
		return nil, errCreateAuthStarredRequest
	}

	return doRequestAndReturnBody(request, g.client)
}

func createRequestWithAuthentication(apiURL, endpoint, token, version string) *http.Request {
	url, err := url.Parse(fmt.Sprintf("%s%s", apiURL, endpoint))
	if err != nil {
		return nil
	}

	headers := map[string][]string{
		versionHeaderName: {version},
		authHeaderName:    {fmt.Sprintf(authHeaderPrefix, token)},
	}

	//nolint:exhaustruct // All the properties do not need to be defined.
	return &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}
}

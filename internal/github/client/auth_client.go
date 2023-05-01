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
		version: defaultApiVersion,
		client:  http.DefaultClient,
	}
}

func (g *GithubClientWithAuthentication) GetUserRepositories(username string) ([]byte, error) {
	request := createRequestWithAuthentication(apiUrl, authListRepositoryEndpoint, g.token, g.version)
	if request == nil {
		return nil, errors.New("could not create authenticated list repository request")
	}
	return doRequestAndReturnBody(request, g.client)
}

func (g *GithubClientWithAuthentication) GetStarredRepositories(username string) ([]byte, error) {
	request := createRequestWithAuthentication(apiUrl, authListStarredEndpoint, g.token, g.version)
	if request == nil {
		return nil, errors.New("could not create authenticated list starred request")
	}
	return doRequestAndReturnBody(request, g.client)
}

func createRequestWithAuthentication(apiUrl, endpoint, token, version string) *http.Request {
	url, err := url.Parse(fmt.Sprintf("%s%s", apiUrl, endpoint))
	if err != nil {
		return nil
	}

	headers := map[string][]string{
		versionHeaderName: {version},
		authHeaderName:    {fmt.Sprintf(authHeaderPrefix, token)},
	}

	return &http.Request{
		Method: http.MethodGet,
		URL:    url,
		Header: headers,
	}
}

package github

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
    "errors"
)

const (
    apiUrl = "https://api.github.com"
    authListRepositoryEndpoint = "/user/repos"
    listRepositoryEndpoint = "/users/%s/repos"
    authListStarredEndpoint = "/user/starred"
    listStarredEndpoint = "/users/%s/starred"

    defaultApiVersion = "2022-11-28"
    
    versionHeaderName = "X-GitHub-Api-Version"
    authHeaderName = "Authorization"
    authHeaderPrefix = "Bearer %s"
)

type GithubNoAuthClient struct {
    version string
    client *http.Client
}

func NewGithubNoAuthClient() *GithubNoAuthClient {
    return &GithubNoAuthClient{
        version: defaultApiVersion,
        client: http.DefaultClient,
    }
}

func NewGithubNoAuthClientWithVersion(version string) *GithubNoAuthClient {
    if version == "" {
        return nil
    }

    return &GithubNoAuthClient{
        version: version,
        client: http.DefaultClient,
    } 
}

func (g *GithubNoAuthClient) GetUserRepositories(username string) ([]byte, error){
    request := createRequest(apiUrl, listRepositoryEndpoint, username, g.version)
    if request == nil {
        return nil, errors.New("could not create authenticated list repository request")
    }
    return doRequestAndReturnBody(request, g.client)
}

func (g *GithubNoAuthClient) GetStarredRepositories(username string) ([]byte, error) {
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
        URL: url,
        Header: headers,
    }
}

func doRequestAndReturnBody(request *http.Request, client *http.Client) ([]byte, error) {
    response, err := client.Do(request)
    if err != nil {
        return nil, fmt.Errorf("failed to do request: %s", err.Error())
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

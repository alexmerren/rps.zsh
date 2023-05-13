package github_test

import (
	"errors"
	"testing"

	"github.com/alexmerren/rps/internal/github"
	"github.com/alexmerren/rps/internal/github/repository"
	"github.com/alexmerren/rps/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	testValidRawRepositoryData = `
    [
        {
            "html_url": "https://github.com/test-user-1/test-repository-1"
        },
        {
            "html_url": "https://github.com/test-user-1/test-repository-2"
        },
        {
            "html_url": "https://github.com/test-user-2/test-repository-3"
        }
    ]`

	testInvalidRawRepositoryData = `
    [
        {
            "not_html_url": "https://github.com/test-user-1/test-repository-1"
        },
    ]`

	testUser1 = "test-user-1"
	testUser2 = "test-user-2"
	testHost  = "github.com"
	testRepo1 = "test-repository-1"
	testRepo2 = "test-repository-2"
	testRepo3 = "test-repository-3"
)

//nolint:gochecknoglobals // What the hell do you want me to do? Not have globals for testing?
var (
	errTest                     = errors.New("test error")
	errParsingJSON              = errors.New("could not format repositories: error in parsing JSON: Unknown value type")
	testProcessedRepositoryData = []*repository.Repository{
		repository.NewRepositoryWithHost(testUser1, testRepo1, testHost),
		repository.NewRepositoryWithHost(testUser1, testRepo2, testHost),
		repository.NewRepositoryWithHost(testUser2, testRepo3, testHost),
	}
)

func TestGetUserRepositories_HappyPath(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetUserRepositories", testUser1).Return([]byte(testValidRawRepositoryData), nil).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetUserRepositories(testUser1)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, testProcessedRepositoryData, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetUserRepositories_ClientErrors(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetUserRepositories", testUser1).Return(nil, errTest).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetUserRepositories(testUser1)

	// assert
	assert.ErrorIs(t, err, errTest)
	assert.Nil(t, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetUserRepositories_ParsingError(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetUserRepositories", testUser1).Return([]byte(testInvalidRawRepositoryData), nil).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetUserRepositories(testUser1)

	// assert
	assert.ErrorContains(t, err, errParsingJSON.Error())
	assert.Nil(t, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetStarredRepositories_ClientError(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetStarredRepositories", testUser1).Return(nil, errTest).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetStarredRepositories(testUser1)

	// assert
	assert.ErrorIs(t, err, errTest)
	assert.Nil(t, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetStarredRepositories_ParsingError(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetStarredRepositories", testUser1).Return([]byte(testInvalidRawRepositoryData), nil).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetStarredRepositories(testUser1)

	// assert
	assert.ErrorContains(t, err, errParsingJSON.Error())
	assert.Nil(t, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetStarredRepositories_HappyPath(t *testing.T) {
	// arrange
	t.Parallel()
	mockClient := mocks.NewGithubInteractor(t)
	mockClient.On("GetStarredRepositories", testUser1).Return([]byte(testValidRawRepositoryData), nil).Once()
	githubUserAPI := github.NewUserAPI(mockClient)

	// act
	repositories, err := githubUserAPI.GetStarredRepositories(testUser1)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, testProcessedRepositoryData, repositories)
	mockClient.AssertExpectations(t)
}

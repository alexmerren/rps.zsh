package github_test

import (
	"context"
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

var (
	testProcessedRepositoryData = []*repository.Repository{
		repository.NewRepositoryWithHost(testUser1, testRepo1, testHost),
		repository.NewRepositoryWithHost(testUser1, testRepo2, testHost),
		repository.NewRepositoryWithHost(testUser2, testRepo3, testHost),
	}
)

func TestGetUserRepositories_HappyPath(t *testing.T) {
	// arrange
	ctx := context.Background()
	mockClient := &mocks.GithubInteractor{}
	mockClient.On("GetUserRepositories", testUser1).Return([]byte(testValidRawRepositoryData), nil).Once()
	githubUserApi := github.NewGithubUserApi(ctx, mockClient)

	// act
	repositories, err := githubUserApi.GetUserRepositories(testUser1)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, testProcessedRepositoryData, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetUserRepositories_Errors(t *testing.T) {
	// arrange
	ctx := context.Background()
	mockClient := &mocks.GithubInteractor{}
    githubUserApi := github.NewGithubUserApi(ctx, mockClient)

	var testCases = []struct {
		name        string
		clientErr   bool
		err         error
	}{
		{
			name:      "Client Error",
			clientErr: true,
			err:       errors.New("test-error"),
		},
		{
			name:       "Reponse Error",
			err:        errors.New("Unknown value type"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.clientErr {
				mockClient.On("GetUserRepositories", testUser1).Return(nil, testCase.err).Once()
			} else {
				mockClient.On("GetUserRepositories", testUser1).Return([]byte(testInvalidRawRepositoryData), nil).Once()
			}
			// act
			repositories, err := githubUserApi.GetUserRepositories(testUser1)

			// assert
			assert.ErrorContains(t, err, testCase.err.Error())
			assert.Nil(t, repositories)
			mockClient.AssertExpectations(t)
		})
	}
}

func TestGetStarredRepositories_HappyPath(t *testing.T) {
	// arrange
	ctx := context.Background()
	mockClient := &mocks.GithubInteractor{}
	mockClient.On("GetStarredRepositories", testUser1).Return([]byte(testValidRawRepositoryData), nil).Once()
	githubUserApi := github.NewGithubUserApi(ctx, mockClient)

	// act
	repositories, err := githubUserApi.GetStarredRepositories(testUser1)

	// assert
	assert.NoError(t, err)
	assert.Equal(t, testProcessedRepositoryData, repositories)
	mockClient.AssertExpectations(t)
}

func TestGetStarredRepositories_Errors(t *testing.T) {
	// arrange
	ctx := context.Background()
	mockClient := &mocks.GithubInteractor{}
    githubUserApi := github.NewGithubUserApi(ctx, mockClient)

	var testCases = []struct {
		name        string
		clientErr   bool
		err         error
	}{
		{
			name:      "Client Error",
			clientErr: true,
			err:       errors.New("test-error"),
		},
		{
			name:       "Reponse Error",
			err:        errors.New("Unknown value type"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.clientErr {
				mockClient.On("GetStarredRepositories", testUser1).Return(nil, testCase.err).Once()
			} else {
				mockClient.On("GetStarredRepositories", testUser1).Return([]byte(testInvalidRawRepositoryData), nil).Once()
			}
			// act
			repositories, err := githubUserApi.GetStarredRepositories(testUser1)

			// assert
			assert.ErrorContains(t, err, testCase.err.Error())
			assert.Nil(t, repositories)
			mockClient.AssertExpectations(t)
		})
	}
}

// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/alexmerren/rps/internal/github/repository"
	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// GetStarredRepositories provides a mock function with given fields: username
func (_m *API) GetStarredRepositories(username string) ([]*repository.Repository, error) {
	ret := _m.Called(username)

	var r0 []*repository.Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*repository.Repository, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) []*repository.Repository); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserRepositories provides a mock function with given fields: username
func (_m *API) GetUserRepositories(username string) ([]*repository.Repository, error) {
	ret := _m.Called(username)

	var r0 []*repository.Repository
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*repository.Repository, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) []*repository.Repository); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repository.Repository)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAPI creates a new instance of API. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *API {
	mock := &API{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
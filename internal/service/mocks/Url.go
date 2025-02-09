// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Url is an autogenerated mock type for the Url type
type Url struct {
	mock.Mock
}

// CreateURL provides a mock function with given fields: urlToSave, alias
func (_m *Url) CreateURL(urlToSave string, alias string) (int, error) {
	ret := _m.Called(urlToSave, alias)

	if len(ret) == 0 {
		panic("no return value specified for CreateURL")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int, error)); ok {
		return rf(urlToSave, alias)
	}
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(urlToSave, alias)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(urlToSave, alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteURLbyID provides a mock function with given fields: id
func (_m *Url) DeleteURLbyID(id int) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteURLbyID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetURL provides a mock function with given fields: alias
func (_m *Url) GetURL(alias string) (string, error) {
	ret := _m.Called(alias)

	if len(ret) == 0 {
		panic("no return value specified for GetURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(alias)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(alias)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(alias)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUrl creates a new instance of Url. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUrl(t interface {
	mock.TestingT
	Cleanup(func())
}) *Url {
	mock := &Url{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

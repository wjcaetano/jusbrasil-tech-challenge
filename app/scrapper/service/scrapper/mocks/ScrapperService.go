// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	scrapper "jusbrasil-tech-challenge/app/scrapper"

	mock "github.com/stretchr/testify/mock"
)

// ScrapperService is an autogenerated mock type for the ScrapperService type
type ScrapperService struct {
	mock.Mock
}

type ScrapperService_Expecter struct {
	mock *mock.Mock
}

func (_m *ScrapperService) EXPECT() *ScrapperService_Expecter {
	return &ScrapperService_Expecter{mock: &_m.Mock}
}

// GetLegalCases provides a mock function with given fields: url
func (_m *ScrapperService) GetLegalCases(url string) ([]scrapper.LegalCase, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for GetLegalCases")
	}

	var r0 []scrapper.LegalCase
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]scrapper.LegalCase, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) []scrapper.LegalCase); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]scrapper.LegalCase)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ScrapperService_GetLegalCases_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLegalCases'
type ScrapperService_GetLegalCases_Call struct {
	*mock.Call
}

// GetLegalCases is a helper method to define mock.On call
//   - url string
func (_e *ScrapperService_Expecter) GetLegalCases(url interface{}) *ScrapperService_GetLegalCases_Call {
	return &ScrapperService_GetLegalCases_Call{Call: _e.mock.On("GetLegalCases", url)}
}

func (_c *ScrapperService_GetLegalCases_Call) Run(run func(url string)) *ScrapperService_GetLegalCases_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ScrapperService_GetLegalCases_Call) Return(_a0 []scrapper.LegalCase, _a1 error) *ScrapperService_GetLegalCases_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ScrapperService_GetLegalCases_Call) RunAndReturn(run func(string) ([]scrapper.LegalCase, error)) *ScrapperService_GetLegalCases_Call {
	_c.Call.Return(run)
	return _c
}

// NewScrapperService creates a new instance of ScrapperService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScrapperService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ScrapperService {
	mock := &ScrapperService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

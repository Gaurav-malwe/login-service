// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Gaurav-malwe/login-service/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// IAuthService is an autogenerated mock type for the IAuthService type
type IAuthService struct {
	mock.Mock
}

// ConfirmUser provides a mock function with given fields: ctx, loginRequest
func (_m *IAuthService) ConfirmUser(ctx context.Context, loginRequest *model.ConfirmRequest) error {
	ret := _m.Called(ctx, loginRequest)

	if len(ret) == 0 {
		panic("no return value specified for ConfirmUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ConfirmRequest) error); ok {
		r0 = rf(ctx, loginRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: ctx, loginRequest
func (_m *IAuthService) LoginUser(ctx context.Context, loginRequest *model.LoginRequest) (string, error) {
	ret := _m.Called(ctx, loginRequest)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.LoginRequest) (string, error)); ok {
		return rf(ctx, loginRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.LoginRequest) string); ok {
		r0 = rf(ctx, loginRequest)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.LoginRequest) error); ok {
		r1 = rf(ctx, loginRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, userRequest
func (_m *IAuthService) RegisterUser(ctx context.Context, userRequest *model.RegisterUserRequest) error {
	ret := _m.Called(ctx, userRequest)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.RegisterUserRequest) error); ok {
		r0 = rf(ctx, userRequest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIAuthService creates a new instance of IAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthService {
	mock := &IAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

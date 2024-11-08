// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// IAuthController is an autogenerated mock type for the IAuthController type
type IAuthController struct {
	mock.Mock
}

// Confirm provides a mock function with given fields: ginCtx
func (_m *IAuthController) Confirm(ginCtx *gin.Context) {
	_m.Called(ginCtx)
}

// Login provides a mock function with given fields: ginCtx
func (_m *IAuthController) Login(ginCtx *gin.Context) {
	_m.Called(ginCtx)
}

// Register provides a mock function with given fields: ginCtx
func (_m *IAuthController) Register(ginCtx *gin.Context) {
	_m.Called(ginCtx)
}

// NewIAuthController creates a new instance of IAuthController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthController(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthController {
	mock := &IAuthController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

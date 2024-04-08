// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/fyfirman/auth-management-go/internal/dto"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceInterface is an autogenerated mock type for the UserServiceInterface type
type UserServiceInterface struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, req
func (_m *UserServiceInterface) Login(ctx context.Context, req dto.LoginRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *dto.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) (*dto.LoginResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, req
func (_m *UserServiceInterface) RegisterUser(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 *dto.RegisterResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.RegisterRequest) (*dto.RegisterResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.RegisterRequest) *dto.RegisterResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RegisterResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.RegisterRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetPassword provides a mock function with given fields: ctx, req
func (_m *UserServiceInterface) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 *dto.ResetPasswordResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.ResetPasswordRequest) (*dto.ResetPasswordResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.ResetPasswordRequest) *dto.ResetPasswordResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResetPasswordResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.ResetPasswordRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserServiceInterface creates a new instance of UserServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceInterface {
	mock := &UserServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
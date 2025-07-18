// Code generated by mockery v2.53.4. DO NOT EDIT.

// Regenerate this file using `make platform-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"

	request "github.com/mattermost/mattermost/server/public/shared/request"
)

// SuiteIFace is an autogenerated mock type for the SuiteIFace type
type SuiteIFace struct {
	mock.Mock
}

// GetSession provides a mock function with given fields: token
func (_m *SuiteIFace) GetSession(token string) (*model.Session, *model.AppError) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for GetSession")
	}

	var r0 *model.Session
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.Session, *model.AppError)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Session); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// HasPermissionToReadChannel provides a mock function with given fields: c, userID, channel
func (_m *SuiteIFace) HasPermissionToReadChannel(c request.CTX, userID string, channel *model.Channel) bool {
	ret := _m.Called(c, userID, channel)

	if len(ret) == 0 {
		panic("no return value specified for HasPermissionToReadChannel")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(request.CTX, string, *model.Channel) bool); ok {
		r0 = rf(c, userID, channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RolesGrantPermission provides a mock function with given fields: roleNames, permissionId
func (_m *SuiteIFace) RolesGrantPermission(roleNames []string, permissionId string) bool {
	ret := _m.Called(roleNames, permissionId)

	if len(ret) == 0 {
		panic("no return value specified for RolesGrantPermission")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, string) bool); ok {
		r0 = rf(roleNames, permissionId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UserCanSeeOtherUser provides a mock function with given fields: c, userID, otherUserId
func (_m *SuiteIFace) UserCanSeeOtherUser(c request.CTX, userID string, otherUserId string) (bool, *model.AppError) {
	ret := _m.Called(c, userID, otherUserId)

	if len(ret) == 0 {
		panic("no return value specified for UserCanSeeOtherUser")
	}

	var r0 bool
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) (bool, *model.AppError)); ok {
		return rf(c, userID, otherUserId)
	}
	if rf, ok := ret.Get(0).(func(request.CTX, string, string) bool); ok {
		r0 = rf(c, userID, otherUserId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(request.CTX, string, string) *model.AppError); ok {
		r1 = rf(c, userID, otherUserId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// NewSuiteIFace creates a new instance of SuiteIFace. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSuiteIFace(t interface {
	mock.TestingT
	Cleanup(func())
}) *SuiteIFace {
	mock := &SuiteIFace{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BlacklistRepo is an autogenerated mock type for the BlacklistRepo type
type BlacklistRepo struct {
	mock.Mock
}

// GetBlacklists provides a mock function with given fields:
func (_m *BlacklistRepo) GetBlacklists() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

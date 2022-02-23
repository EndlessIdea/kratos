// Code generated by mockery v2.9.4. DO NOT EDIT.

package kafka

import mock "github.com/stretchr/testify/mock"

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: message
func (_m *MockHandler) Handle(message Message) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(Message) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topic provides a mock function with given fields:
func (_m *MockHandler) Topic() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

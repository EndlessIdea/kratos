// Code generated by mockery v2.9.4. DO NOT EDIT.

package kafka

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockConsumer is an autogenerated mock type for the Consumer type
type MockConsumer struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockConsumer) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Consume provides a mock function with given fields: ctx
func (_m *MockConsumer) Consume(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterHandler provides a mock function with given fields: handler
func (_m *MockConsumer) RegisterHandler(handler Handler) {
	_m.Called(handler)
}

// Topic provides a mock function with given fields:
func (_m *MockConsumer) Topic() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

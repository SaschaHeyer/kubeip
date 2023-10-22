// Code generated by mockery v2.30.16. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// EipAssigner is an autogenerated mock type for the EipAssigner type
type EipAssigner struct {
	mock.Mock
}

type EipAssigner_Expecter struct {
	mock *mock.Mock
}

func (_m *EipAssigner) EXPECT() *EipAssigner_Expecter {
	return &EipAssigner_Expecter{mock: &_m.Mock}
}

// Assign provides a mock function with given fields: ctx, networkInterfaceID, allocationID
func (_m *EipAssigner) Assign(ctx context.Context, networkInterfaceID string, allocationID string) error {
	ret := _m.Called(ctx, networkInterfaceID, allocationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, networkInterfaceID, allocationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EipAssigner_Assign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Assign'
type EipAssigner_Assign_Call struct {
	*mock.Call
}

// Assign is a helper method to define mock.On call
//   - ctx context.Context
//   - networkInterfaceID string
//   - allocationID string
func (_e *EipAssigner_Expecter) Assign(ctx interface{}, networkInterfaceID interface{}, allocationID interface{}) *EipAssigner_Assign_Call {
	return &EipAssigner_Assign_Call{Call: _e.mock.On("Assign", ctx, networkInterfaceID, allocationID)}
}

func (_c *EipAssigner_Assign_Call) Run(run func(ctx context.Context, networkInterfaceID string, allocationID string)) *EipAssigner_Assign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *EipAssigner_Assign_Call) Return(_a0 error) *EipAssigner_Assign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EipAssigner_Assign_Call) RunAndReturn(run func(context.Context, string, string) error) *EipAssigner_Assign_Call {
	_c.Call.Return(run)
	return _c
}

// Unassign provides a mock function with given fields: ctx, associationID
func (_m *EipAssigner) Unassign(ctx context.Context, associationID string) error {
	ret := _m.Called(ctx, associationID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, associationID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EipAssigner_Unassign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Unassign'
type EipAssigner_Unassign_Call struct {
	*mock.Call
}

// Unassign is a helper method to define mock.On call
//   - ctx context.Context
//   - associationID string
func (_e *EipAssigner_Expecter) Unassign(ctx interface{}, associationID interface{}) *EipAssigner_Unassign_Call {
	return &EipAssigner_Unassign_Call{Call: _e.mock.On("Unassign", ctx, associationID)}
}

func (_c *EipAssigner_Unassign_Call) Run(run func(ctx context.Context, associationID string)) *EipAssigner_Unassign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *EipAssigner_Unassign_Call) Return(_a0 error) *EipAssigner_Unassign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EipAssigner_Unassign_Call) RunAndReturn(run func(context.Context, string) error) *EipAssigner_Unassign_Call {
	_c.Call.Return(run)
	return _c
}

// NewEipAssigner creates a new instance of EipAssigner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEipAssigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *EipAssigner {
	mock := &EipAssigner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

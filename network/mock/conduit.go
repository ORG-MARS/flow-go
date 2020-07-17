// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/dapperlabs/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Conduit is an autogenerated mock type for the Conduit type
type Conduit struct {
	mock.Mock
}

// Publish provides a mock function with given fields: message, restrictor
func (_m *Conduit) Publish(message interface{}, restrictor flow.IdentityFilter) error {
	ret := _m.Called(message, restrictor)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, flow.IdentityFilter) error); ok {
		r0 = rf(message, restrictor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: message, num, selector
func (_m *Conduit) Send(message interface{}, num uint, selector flow.IdentityFilter) error {
	ret := _m.Called(message, num, selector)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, uint, flow.IdentityFilter) error); ok {
		r0 = rf(message, num, selector)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Transmit provides a mock function with given fields: message, recipientIDs
func (_m *Conduit) Transmit(message interface{}, recipientIDs ...flow.Identifier) error {
	_va := make([]interface{}, len(recipientIDs))
	for _i := range recipientIDs {
		_va[_i] = recipientIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...flow.Identifier) error); ok {
		r0 = rf(message, recipientIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

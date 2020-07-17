// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import (
	flow "github.com/dapperlabs/flow-go/model/flow"
	message "github.com/dapperlabs/flow-go/network/gossip/libp2p/message"
	middleware "github.com/dapperlabs/flow-go/network/gossip/libp2p/middleware"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// Publish provides a mock function with given fields: msg
func (_m *Middleware) Publish(msg *message.Message) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Message) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: msg, recipientIDs
func (_m *Middleware) Send(msg *message.Message, recipientIDs ...flow.Identifier) error {
	_va := make([]interface{}, len(recipientIDs))
	for _i := range recipientIDs {
		_va[_i] = recipientIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(*message.Message, ...flow.Identifier) error); ok {
		r0 = rf(msg, recipientIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: overlay
func (_m *Middleware) Start(overlay middleware.Overlay) error {
	ret := _m.Called(overlay)

	var r0 error
	if rf, ok := ret.Get(0).(func(middleware.Overlay) error); ok {
		r0 = rf(overlay)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Middleware) Stop() {
	_m.Called()
}

// Subscribe provides a mock function with given fields: channelID
func (_m *Middleware) Subscribe(channelID uint8) error {
	ret := _m.Called(channelID)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint8) error); ok {
		r0 = rf(channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mock

import crypto "github.com/dapperlabs/flow-go/crypto"
import flow "github.com/dapperlabs/flow-go/model/flow"
import mock "github.com/stretchr/testify/mock"

// HotStuff is an autogenerated mock type for the HotStuff type
type HotStuff struct {
	mock.Mock
}

// Start provides a mock function with given fields:
func (_m *HotStuff) Start() (func(), <-chan struct{}) {
	ret := _m.Called()

	var r0 func()
	if rf, ok := ret.Get(0).(func() func()); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	var r1 <-chan struct{}
	if rf, ok := ret.Get(1).(func() <-chan struct{}); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan struct{})
		}
	}

	return r0, r1
}

// SubmitProposal provides a mock function with given fields: proposal, parentView
func (_m *HotStuff) SubmitProposal(proposal *flow.Header, parentView uint64) {
	_m.Called(proposal, parentView)
}

// SubmitVote provides a mock function with given fields: originID, blockID, view, sig
func (_m *HotStuff) SubmitVote(originID flow.Identifier, blockID flow.Identifier, view uint64, sig crypto.Signature) {
	_m.Called(originID, blockID, view, sig)
}

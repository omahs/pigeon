// Code generated by mockery v2.11.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	cronchain "github.com/palomachain/sparrow/types/cronchain"

	testing "testing"
)

// QueryValsetServer is an autogenerated mock type for the QueryValsetServer type
type QueryValsetServer struct {
	mock.Mock
}

// ValidatorInfo provides a mock function with given fields: _a0, _a1
func (_m *QueryValsetServer) ValidatorInfo(_a0 context.Context, _a1 *cronchain.QueryValidatorInfoRequest) (*cronchain.QueryValidatorInfoResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *cronchain.QueryValidatorInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context, *cronchain.QueryValidatorInfoRequest) *cronchain.QueryValidatorInfoResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cronchain.QueryValidatorInfoResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *cronchain.QueryValidatorInfoRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryValsetServer creates a new instance of QueryValsetServer. It also registers a cleanup function to assert the mocks expectations.
func NewQueryValsetServer(t testing.TB) *QueryValsetServer {
	mock := &QueryValsetServer{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

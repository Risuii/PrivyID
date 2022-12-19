package mocks

import (
	"context"
	"privyID/helpers/response"
	"privyID/models"

	"github.com/stretchr/testify/mock"
)

type MockCake struct {
	mock.Mock
}

func (_m *MockCake) AddCakes(ctx context.Context, params models.CheeseCake) response.Response {
	ret := _m.Called(ctx, params)

	var r0 response.Response

	if rf, ok := ret.Get(0).(func(context.Context, models.CheeseCake) response.Response); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}
	return r0
}

func (_m *MockCake) DetailCakes(ctx context.Context, id int64) response.Response {
	ret := _m.Called(ctx, id)

	var r0 response.Response

	if rf, ok := ret.Get(0).(func(context.Context, int64) response.Response); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

func (_m *MockCake) ListCakes() response.Response {
	ret := _m.Called()

	var r0 response.Response
	if rf, ok := ret.Get(0).(func() response.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

func (_m *MockCake) UpdateCake(ctx context.Context, id int64, params models.CheeseCake) response.Response {
	ret := _m.Called(ctx, id, params)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64, models.CheeseCake) response.Response); ok {
		r0 = rf(ctx, id, params)
	}

	return r0
}

func (_m *MockCake) DeleteCake(ctx context.Context, id int64) response.Response {
	ret := _m.Called(ctx, id)

	var r0 response.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64) response.Response); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(response.Response)
		}
	}

	return r0
}

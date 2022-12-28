package mocks

import (
	"context"
	"privyID/models"

	mock "github.com/stretchr/testify/mock"
)

type CakeRepository struct {
	mock.Mock
}

func (_m *CakeRepository) Create(ctx context.Context, params models.CheeseCake) (int64, error) {
	ret := _m.Called(ctx, params)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, models.CheeseCake) int64); ok {
		r0 = rf(ctx, params)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.CheeseCake) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CakeRepository) FindByID(ctx context.Context, ID int64) (models.CheeseCake, error) {
	ret := _m.Called(ctx, ID)

	var r0 models.CheeseCake
	if rf, ok := ret.Get(0).(func(context.Context, int64) models.CheeseCake); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(models.CheeseCake)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CakeRepository) FindAll() ([]models.CheeseCake, error) {
	ret := _m.Called()

	var r0 []models.CheeseCake
	if rf, ok := ret.Get(0).(func() []models.CheeseCake); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.CheeseCake)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *CakeRepository) Update(ctx context.Context, id int64, _a2 models.CheeseCake) error {
	ret := _m.Called(ctx, id, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, models.CheeseCake) error); ok {
		r0 = rf(ctx, id, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *CakeRepository) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

package mocks

import helpers "github.com/ghmeier/bloodlines/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/ghmeier/bloodlines/models"
import uuid "github.com/pborman/uuid"

// ReceiptI is an autogenerated mock type for the ReceiptI type
type ReceiptI struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *ReceiptI) GetAll(_a0 int, _a1 int) ([]*models.Receipt, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.Receipt
	if rf, ok := ret.Get(0).(func(int, int) []*models.Receipt); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0
func (_m *ReceiptI) GetByID(_a0 string) (*models.Receipt, error) {
	ret := _m.Called(_a0)

	var r0 *models.Receipt
	if rf, ok := ret.Get(0).(func(string) *models.Receipt); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *ReceiptI) Insert(_a0 *models.Receipt) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Receipt) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetStatus provides a mock function with given fields: _a0, _a1
func (_m *ReceiptI) SetStatus(_a0 uuid.UUID, _a1 models.Status) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, models.Status) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ helpers.ReceiptI = (*ReceiptI)(nil)
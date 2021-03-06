package mocks

import helpers "github.com/ghmeier/bloodlines/helpers"
import mock "github.com/stretchr/testify/mock"
import models "github.com/ghmeier/bloodlines/models"

// PreferenceI is an autogenerated mock type for the PreferenceI type
type PreferenceI struct {
	mock.Mock
}

// GetAll provides a mock function with given fields: _a0, _a1
func (_m *PreferenceI) GetAll(_a0 int, _a1 int) ([]*models.Preference, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.Preference
	if rf, ok := ret.Get(0).(func(int, int) []*models.Preference); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Preference)
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

// GetByUserID provides a mock function with given fields: _a0
func (_m *PreferenceI) GetByUserID(_a0 string) (*models.Preference, error) {
	ret := _m.Called(_a0)

	var r0 *models.Preference
	if rf, ok := ret.Get(0).(func(string) *models.Preference); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Preference)
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
func (_m *PreferenceI) Insert(_a0 *models.Preference) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Preference) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *PreferenceI) Update(_a0 *models.Preference) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Preference) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ helpers.PreferenceI = (*PreferenceI)(nil)

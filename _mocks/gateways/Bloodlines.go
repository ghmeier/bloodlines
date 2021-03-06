package mocks

import gateways "github.com/ghmeier/bloodlines/gateways"
import mock "github.com/stretchr/testify/mock"
import models "github.com/ghmeier/bloodlines/models"
import uuid "github.com/pborman/uuid"

// Bloodlines is an autogenerated mock type for the Bloodlines type
type Bloodlines struct {
	mock.Mock
}

// ActivateTrigger provides a mock function with given fields: key, receipt
func (_m *Bloodlines) ActivateTrigger(key string, receipt *models.Receipt) (*models.SendRequest, error) {
	ret := _m.Called(key, receipt)

	var r0 *models.SendRequest
	if rf, ok := ret.Get(0).(func(string, *models.Receipt) *models.SendRequest); ok {
		r0 = rf(key, receipt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.SendRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *models.Receipt) error); ok {
		r1 = rf(key, receipt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContent provides a mock function with given fields: id
func (_m *Bloodlines) DeleteContent(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletePreference provides a mock function with given fields: id
func (_m *Bloodlines) DeletePreference(id uuid.UUID) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTrigger provides a mock function with given fields: key
func (_m *Bloodlines) DeleteTrigger(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllContent provides a mock function with given fields: offset, limit
func (_m *Bloodlines) GetAllContent(offset int, limit int) ([]*models.Content, error) {
	ret := _m.Called(offset, limit)

	var r0 []*models.Content
	if rf, ok := ret.Get(0).(func(int, int) []*models.Content); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Content)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReceipts provides a mock function with given fields: offset, limit
func (_m *Bloodlines) GetAllReceipts(offset int, limit int) ([]*models.Receipt, error) {
	ret := _m.Called(offset, limit)

	var r0 []*models.Receipt
	if rf, ok := ret.Get(0).(func(int, int) []*models.Receipt); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTriggers provides a mock function with given fields: offset, limit
func (_m *Bloodlines) GetAllTriggers(offset int, limit int) ([]*models.Trigger, error) {
	ret := _m.Called(offset, limit)

	var r0 []*models.Trigger
	if rf, ok := ret.Get(0).(func(int, int) []*models.Trigger); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContentByID provides a mock function with given fields: id
func (_m *Bloodlines) GetContentByID(id uuid.UUID) (*models.Content, error) {
	ret := _m.Called(id)

	var r0 *models.Content
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Content); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Content)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPreference provides a mock function with given fields: id
func (_m *Bloodlines) GetPreference(id uuid.UUID) (*models.Preference, error) {
	ret := _m.Called(id)

	var r0 *models.Preference
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Preference); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Preference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReceiptByID provides a mock function with given fields: id
func (_m *Bloodlines) GetReceiptByID(id uuid.UUID) (*models.Receipt, error) {
	ret := _m.Called(id)

	var r0 *models.Receipt
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Receipt); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTriggerByKey provides a mock function with given fields: key
func (_m *Bloodlines) GetTriggerByKey(key string) (*models.Trigger, error) {
	ret := _m.Called(key)

	var r0 *models.Trigger
	if rf, ok := ret.Get(0).(func(string) *models.Trigger); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewContent provides a mock function with given fields: newContent
func (_m *Bloodlines) NewContent(newContent *models.Content) (*models.Content, error) {
	ret := _m.Called(newContent)

	var r0 *models.Content
	if rf, ok := ret.Get(0).(func(*models.Content) *models.Content); ok {
		r0 = rf(newContent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Content)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Content) error); ok {
		r1 = rf(newContent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPreference provides a mock function with given fields: id
func (_m *Bloodlines) NewPreference(id uuid.UUID) (*models.Preference, error) {
	ret := _m.Called(id)

	var r0 *models.Preference
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.Preference); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Preference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTrigger provides a mock function with given fields: t
func (_m *Bloodlines) NewTrigger(t *models.Trigger) (*models.Trigger, error) {
	ret := _m.Called(t)

	var r0 *models.Trigger
	if rf, ok := ret.Get(0).(func(*models.Trigger) *models.Trigger); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Trigger) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendReceipt provides a mock function with given fields: receipt
func (_m *Bloodlines) SendReceipt(receipt *models.Receipt) (*models.Receipt, error) {
	ret := _m.Called(receipt)

	var r0 *models.Receipt
	if rf, ok := ret.Get(0).(func(*models.Receipt) *models.Receipt); ok {
		r0 = rf(receipt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Receipt) error); ok {
		r1 = rf(receipt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContent provides a mock function with given fields: update
func (_m *Bloodlines) UpdateContent(update *models.Content) (*models.Content, error) {
	ret := _m.Called(update)

	var r0 *models.Content
	if rf, ok := ret.Get(0).(func(*models.Content) *models.Content); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Content)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Content) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePreference provides a mock function with given fields: _a0
func (_m *Bloodlines) UpdatePreference(_a0 *models.Preference) (*models.Preference, error) {
	ret := _m.Called(_a0)

	var r0 *models.Preference
	if rf, ok := ret.Get(0).(func(*models.Preference) *models.Preference); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Preference)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Preference) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTrigger provides a mock function with given fields: update
func (_m *Bloodlines) UpdateTrigger(update *models.Trigger) (*models.Trigger, error) {
	ret := _m.Called(update)

	var r0 *models.Trigger
	if rf, ok := ret.Get(0).(func(*models.Trigger) *models.Trigger); ok {
		r0 = rf(update)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Trigger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.Trigger) error); ok {
		r1 = rf(update)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

var _ gateways.Bloodlines = (*Bloodlines)(nil)

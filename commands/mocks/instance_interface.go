package mocks

import "github.com/stretchr/testify/mock"

import swaggerclient "github.com/go-swagger/go-swagger/client"

import "github.com/hpcloud/cf-plugin-usb/lib/models"

type InstanceInterface struct {
	mock.Mock
}

func (_m *InstanceInterface) Create(_a0 swaggerclient.AuthInfoWriter, _a1 []string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(swaggerclient.AuthInfoWriter, []string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(swaggerclient.AuthInfoWriter, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *InstanceInterface) Delete(_a0 swaggerclient.AuthInfoWriter, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(swaggerclient.AuthInfoWriter, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(swaggerclient.AuthInfoWriter, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *InstanceInterface) Update(_a0 swaggerclient.AuthInfoWriter, _a1 []string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(swaggerclient.AuthInfoWriter, []string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(swaggerclient.AuthInfoWriter, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *InstanceInterface) List(_a0 swaggerclient.AuthInfoWriter, _a1 string) ([]*models.DriverInstance, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*models.DriverInstance
	if rf, ok := ret.Get(0).(func(swaggerclient.AuthInfoWriter, string) []*models.DriverInstance); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.DriverInstance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(swaggerclient.AuthInfoWriter, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *InstanceInterface) GetDriverInstanceByName(_a0 swaggerclient.AuthInfoWriter, _a1 string) *models.DriverInstance {
	ret := _m.Called(_a0, _a1)

	var r0 *models.DriverInstance
	if rf, ok := ret.Get(0).(func(swaggerclient.AuthInfoWriter, string) *models.DriverInstance); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.DriverInstance)
		}
	}

	return r0
}

package mocks

import "github.com/stretchr/testify/mock"

import client "github.com/go-swagger/go-swagger/client"
import operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"

type UsbClientInterface struct {
	mock.Mock
}

func (_m *UsbClientInterface) CreateDial(params *operations.CreateDialParams, authInfo client.AuthInfoWriter) (*operations.CreateDialCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.CreateDialCreated
	if rf, ok := ret.Get(0).(func(*operations.CreateDialParams, client.AuthInfoWriter) *operations.CreateDialCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.CreateDialCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.CreateDialParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) CreateDriver(params *operations.CreateDriverParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.CreateDriverCreated
	if rf, ok := ret.Get(0).(func(*operations.CreateDriverParams, client.AuthInfoWriter) *operations.CreateDriverCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.CreateDriverCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.CreateDriverParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) CreateDriverInstance(params *operations.CreateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverInstanceCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.CreateDriverInstanceCreated
	if rf, ok := ret.Get(0).(func(*operations.CreateDriverInstanceParams, client.AuthInfoWriter) *operations.CreateDriverInstanceCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.CreateDriverInstanceCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.CreateDriverInstanceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) DeleteDial(params *operations.DeleteDialParams, authInfo client.AuthInfoWriter) (*operations.DeleteDialNoContent, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.DeleteDialNoContent
	if rf, ok := ret.Get(0).(func(*operations.DeleteDialParams, client.AuthInfoWriter) *operations.DeleteDialNoContent); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.DeleteDialNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.DeleteDialParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) DeleteDriver(params *operations.DeleteDriverParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverNoContent, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.DeleteDriverNoContent
	if rf, ok := ret.Get(0).(func(*operations.DeleteDriverParams, client.AuthInfoWriter) *operations.DeleteDriverNoContent); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.DeleteDriverNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.DeleteDriverParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) DeleteDriverInstance(params *operations.DeleteDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverInstanceNoContent, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.DeleteDriverInstanceNoContent
	if rf, ok := ret.Get(0).(func(*operations.DeleteDriverInstanceParams, client.AuthInfoWriter) *operations.DeleteDriverInstanceNoContent); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.DeleteDriverInstanceNoContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.DeleteDriverInstanceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetAllDials(params *operations.GetAllDialsParams, authInfo client.AuthInfoWriter) (*operations.GetAllDialsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetAllDialsOK
	if rf, ok := ret.Get(0).(func(*operations.GetAllDialsParams, client.AuthInfoWriter) *operations.GetAllDialsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetAllDialsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetAllDialsParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDial(params *operations.GetDialParams, authInfo client.AuthInfoWriter) (*operations.GetDialOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDialOK
	if rf, ok := ret.Get(0).(func(*operations.GetDialParams, client.AuthInfoWriter) *operations.GetDialOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDialOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDialParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDialSchema(params *operations.GetDialSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDialSchemaOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDialSchemaOK
	if rf, ok := ret.Get(0).(func(*operations.GetDialSchemaParams, client.AuthInfoWriter) *operations.GetDialSchemaOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDialSchemaOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDialSchemaParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDriver(params *operations.GetDriverParams, authInfo client.AuthInfoWriter) (*operations.GetDriverOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDriverOK
	if rf, ok := ret.Get(0).(func(*operations.GetDriverParams, client.AuthInfoWriter) *operations.GetDriverOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDriverOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDriverParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDriverInstance(params *operations.GetDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstanceOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDriverInstanceOK
	if rf, ok := ret.Get(0).(func(*operations.GetDriverInstanceParams, client.AuthInfoWriter) *operations.GetDriverInstanceOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDriverInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDriverInstanceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDriverInstances(params *operations.GetDriverInstancesParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstancesOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDriverInstancesOK
	if rf, ok := ret.Get(0).(func(*operations.GetDriverInstancesParams, client.AuthInfoWriter) *operations.GetDriverInstancesOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDriverInstancesOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDriverInstancesParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDriverSchema(params *operations.GetDriverSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDriverSchemaOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDriverSchemaOK
	if rf, ok := ret.Get(0).(func(*operations.GetDriverSchemaParams, client.AuthInfoWriter) *operations.GetDriverSchemaOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDriverSchemaOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDriverSchemaParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetDrivers(params *operations.GetDriversParams, authInfo client.AuthInfoWriter) (*operations.GetDriversOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetDriversOK
	if rf, ok := ret.Get(0).(func(*operations.GetDriversParams, client.AuthInfoWriter) *operations.GetDriversOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetDriversOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetDriversParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetInfo(params *operations.GetInfoParams, authInfo client.AuthInfoWriter) (*operations.GetInfoOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetInfoOK
	if rf, ok := ret.Get(0).(func(*operations.GetInfoParams, client.AuthInfoWriter) *operations.GetInfoOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetInfoOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetInfoParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetService(params *operations.GetServiceParams, authInfo client.AuthInfoWriter) (*operations.GetServiceOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetServiceOK
	if rf, ok := ret.Get(0).(func(*operations.GetServiceParams, client.AuthInfoWriter) *operations.GetServiceOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetServiceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetServiceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo client.AuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetServiceByInstanceIDOK
	if rf, ok := ret.Get(0).(func(*operations.GetServiceByInstanceIDParams, client.AuthInfoWriter) *operations.GetServiceByInstanceIDOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetServiceByInstanceIDOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetServiceByInstanceIDParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetServicePlan(params *operations.GetServicePlanParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlanOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetServicePlanOK
	if rf, ok := ret.Get(0).(func(*operations.GetServicePlanParams, client.AuthInfoWriter) *operations.GetServicePlanOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetServicePlanOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetServicePlanParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) GetServicePlans(params *operations.GetServicePlansParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlansOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.GetServicePlansOK
	if rf, ok := ret.Get(0).(func(*operations.GetServicePlansParams, client.AuthInfoWriter) *operations.GetServicePlansOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.GetServicePlansOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.GetServicePlansParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) PingDriverInstance(params *operations.PingDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.PingDriverInstanceOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.PingDriverInstanceOK
	if rf, ok := ret.Get(0).(func(*operations.PingDriverInstanceParams, client.AuthInfoWriter) *operations.PingDriverInstanceOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.PingDriverInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.PingDriverInstanceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateCatalog(params *operations.UpdateCatalogParams, authInfo client.AuthInfoWriter) (*operations.UpdateCatalogOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateCatalogOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateCatalogParams, client.AuthInfoWriter) *operations.UpdateCatalogOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateCatalogOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateCatalogParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateDial(params *operations.UpdateDialParams, authInfo client.AuthInfoWriter) (*operations.UpdateDialOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateDialOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateDialParams, client.AuthInfoWriter) *operations.UpdateDialOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateDialOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateDialParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateDriver(params *operations.UpdateDriverParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateDriverOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateDriverParams, client.AuthInfoWriter) *operations.UpdateDriverOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateDriverOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateDriverParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateDriverInstance(params *operations.UpdateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverInstanceOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateDriverInstanceOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateDriverInstanceParams, client.AuthInfoWriter) *operations.UpdateDriverInstanceOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateDriverInstanceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateDriverInstanceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateService(params *operations.UpdateServiceParams, authInfo client.AuthInfoWriter) (*operations.UpdateServiceOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateServiceOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateServiceParams, client.AuthInfoWriter) *operations.UpdateServiceOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateServiceOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateServiceParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo client.AuthInfoWriter) (*operations.UpdateServicePlanOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UpdateServicePlanOK
	if rf, ok := ret.Get(0).(func(*operations.UpdateServicePlanParams, client.AuthInfoWriter) *operations.UpdateServicePlanOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UpdateServicePlanOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UpdateServicePlanParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) UploadDriver(params *operations.UploadDriverParams, authInfo client.AuthInfoWriter) (*operations.UploadDriverOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *operations.UploadDriverOK
	if rf, ok := ret.Get(0).(func(*operations.UploadDriverParams, client.AuthInfoWriter) *operations.UploadDriverOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*operations.UploadDriverOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*operations.UploadDriverParams, client.AuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *UsbClientInterface) SetTransport(transport client.Transport) {
	_m.Called(transport)
}

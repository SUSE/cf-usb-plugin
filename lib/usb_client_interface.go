package lib

import (
	client "github.com/go-swagger/go-swagger/client"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
)

//Usb client interface
type UsbClientInterface interface {
	CreateDial(params *operations.CreateDialParams, authInfo client.AuthInfoWriter) (*operations.CreateDialCreated, error)
	CreateDriver(params *operations.CreateDriverParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverCreated, error)
	CreateDriverInstance(params *operations.CreateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverInstanceCreated, error)
	DeleteDial(params *operations.DeleteDialParams, authInfo client.AuthInfoWriter) (*operations.DeleteDialNoContent, error)
	DeleteDriver(params *operations.DeleteDriverParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverNoContent, error)
	DeleteDriverInstance(params *operations.DeleteDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverInstanceNoContent, error)
	GetAllDials(params *operations.GetAllDialsParams, authInfo client.AuthInfoWriter) (*operations.GetAllDialsOK, error)
	GetDial(params *operations.GetDialParams, authInfo client.AuthInfoWriter) (*operations.GetDialOK, error)
	GetDialSchema(params *operations.GetDialSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDialSchemaOK, error)
	GetDriver(params *operations.GetDriverParams, authInfo client.AuthInfoWriter) (*operations.GetDriverOK, error)
	GetDriverInstance(params *operations.GetDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstanceOK, error)
	GetDriverInstances(params *operations.GetDriverInstancesParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstancesOK, error)
	GetDriverSchema(params *operations.GetDriverSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDriverSchemaOK, error)
	GetDrivers(params *operations.GetDriversParams, authInfo client.AuthInfoWriter) (*operations.GetDriversOK, error)
	GetInfo(params *operations.GetInfoParams, authInfo client.AuthInfoWriter) (*operations.GetInfoOK, error)
	GetService(params *operations.GetServiceParams, authInfo client.AuthInfoWriter) (*operations.GetServiceOK, error)
	GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo client.AuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error)
	GetServicePlan(params *operations.GetServicePlanParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlanOK, error)
	GetServicePlans(params *operations.GetServicePlansParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlansOK, error)
	PingDriverInstance(params *operations.PingDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.PingDriverInstanceOK, error)
	UpdateCatalog(params *operations.UpdateCatalogParams, authInfo client.AuthInfoWriter) (*operations.UpdateCatalogOK, error)
	UpdateDial(params *operations.UpdateDialParams, authInfo client.AuthInfoWriter) (*operations.UpdateDialOK, error)
	UpdateDriver(params *operations.UpdateDriverParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverOK, error)
	UpdateDriverInstance(params *operations.UpdateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverInstanceOK, error)
	UpdateService(params *operations.UpdateServiceParams, authInfo client.AuthInfoWriter) (*operations.UpdateServiceOK, error)
	UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo client.AuthInfoWriter) (*operations.UpdateServicePlanOK, error)
	UploadDriver(params *operations.UploadDriverParams, authInfo client.AuthInfoWriter) (*operations.UploadDriverOK, error)
	SetTransport(transport client.Transport)
}

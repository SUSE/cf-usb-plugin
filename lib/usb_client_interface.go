package lib

import (
	"github.com/go-openapi/runtime"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//Usb client interface
type UsbClientInterface interface {
	CreateDial(params *operations.CreateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDialCreated, error)
	CreateDriver(params *operations.CreateDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDriverCreated, error)
	CreateDriverInstance(params *operations.CreateDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDriverInstanceCreated, error)
	DeleteDial(params *operations.DeleteDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDialNoContent, error)
	DeleteDriver(params *operations.DeleteDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDriverNoContent, error)
	DeleteDriverInstance(params *operations.DeleteDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDriverInstanceNoContent, error)
	GetAllDials(params *operations.GetAllDialsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetAllDialsOK, error)
	GetDial(params *operations.GetDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDialOK, error)
	GetDialSchema(params *operations.GetDialSchemaParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDialSchemaOK, error)
	GetDriver(params *operations.GetDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverOK, error)
	GetDriverInstance(params *operations.GetDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverInstanceOK, error)
	GetDriverInstances(params *operations.GetDriverInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverInstancesOK, error)
	GetDriverSchema(params *operations.GetDriverSchemaParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverSchemaOK, error)
	GetDrivers(params *operations.GetDriversParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriversOK, error)
	GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error)
	GetService(params *operations.GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceOK, error)
	GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error)
	GetServicePlan(params *operations.GetServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlanOK, error)
	GetServicePlans(params *operations.GetServicePlansParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlansOK, error)
	PingDriverInstance(params *operations.PingDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverInstanceOK, error)
	UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error)
	UpdateDial(params *operations.UpdateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDialOK, error)
	UpdateDriver(params *operations.UpdateDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverOK, error)
	UpdateDriverInstance(params *operations.UpdateDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverInstanceOK, error)
	UpdateService(params *operations.UpdateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServiceOK, error)
	UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServicePlanOK, error)
	UploadDriver(params *operations.UploadDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UploadDriverOK, error)
	SetTransport(transport runtime.ClientTransport)
	GetDriverByName(authInfo runtime.ClientAuthInfoWriter, driverName string) (*models.Driver, error)
	GetDriverInstanceByName(authHeader runtime.ClientAuthInfoWriter, driverInstanceName string) (*models.DriverInstance, error)
	GetServiceByDriverInstanceID(authInfo runtime.ClientAuthInfoWriter, driverInstanceID string) (*models.Service, error)
	GetPlanByID(authInfo runtime.ClientAuthInfoWriter, planID string) (*models.Plan, error)
}

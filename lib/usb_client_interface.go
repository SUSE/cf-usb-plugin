package lib

import (
	"github.com/go-openapi/runtime"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//Usb client interface
type UsbClientInterface interface {
	CreateDial(params *operations.CreateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDialCreated, error)
	CreateInstance(params *operations.CreateInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateInstanceCreated, error)
	DeleteDial(params *operations.DeleteDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDialNoContent, error)
	DeleteInstance(params *operations.DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteInstanceNoContent, error)
	GetAllDials(params *operations.GetAllDialsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetAllDialsOK, error)
	GetDial(params *operations.GetDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDialOK, error)
	GetInstance(params *operations.GetInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInstanceOK, error)
	GetInstances(params *operations.GetInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInstancesOK, error)
	GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error)
	GetService(params *operations.GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceOK, error)
	GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error)
	GetServicePlan(params *operations.GetServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlanOK, error)
	GetServicePlans(params *operations.GetServicePlansParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlansOK, error)
	PingInstance(params *operations.PingInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingInstanceOK, error)
	UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error)
	UpdateDial(params *operations.UpdateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDialOK, error)
	UpdateInstance(params *operations.UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateInstanceOK, error)
	UpdateService(params *operations.UpdateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServiceOK, error)
	UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServicePlanOK, error)
	SetTransport(transport runtime.ClientTransport)
	GetInstanceByName(authHeader runtime.ClientAuthInfoWriter, driverInstanceName string) (*models.Instance, error)
	GetServiceByDriverInstanceID(authInfo runtime.ClientAuthInfoWriter, driverInstanceID string) (*models.Service, error)
	GetPlanByID(authInfo runtime.ClientAuthInfoWriter, planID string) (*models.Plan, error)
}

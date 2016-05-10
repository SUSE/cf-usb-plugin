package lib

import (
	"github.com/go-openapi/runtime"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//Usb client interface
type UsbClientInterface interface {
	RegisterDriverEndpoint(params *operations.RegisterDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.RegisterDriverEndpointCreated, error)
	UnregisterDriverEndpoint(params *operations.UnregisterDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UnregisterDriverInstanceNoContent, error)
	GetDriverEndpoint(params *operations.GetDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointOK, error)
	GetDriverEndpointByName(instanceName string, authInfo runtime.ClientAuthInfoWriter) (*models.DriverEndpoint, error)
	GetDriverEndpoints(params *operations.GetDriverEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointsOK, error)
	GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error)
	PingEndpoint(params *operations.PingDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverEndpointOK, error)
	UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error)
	UpdateDriverEndpoint(params *operations.UpdateDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverEndpointOK, error)
	SetTransport(transport runtime.ClientTransport)
}

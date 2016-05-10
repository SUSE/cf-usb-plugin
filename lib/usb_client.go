package lib

import (
	"net/url"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type UsbClient struct {
	httpClient *operations.Client
}

func NewUsbClient(target *url.URL, trace bool) UsbClientInterface {
	transport := httptransport.New(target.Host, "/", []string{target.Scheme})
	transport.Debug = trace
	return &UsbClient{httpClient: operations.New(transport, strfmt.Default)}
}

func (a *UsbClient) RegisterDriverEndpoint(params *operations.RegisterDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.RegisterDriverEndpointCreated, error) {
	response, err := a.httpClient.RegisterDriverEndpoint(params, authInfo)
	return response, err
}

func (a *UsbClient) UnregisterDriverEndpoint(params *operations.UnregisterDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UnregisterDriverInstanceNoContent, error) {
	response, err := a.httpClient.UnregisterDriverInstance(params, authInfo)
	return response, err
}

func (a *UsbClient) GetDriverEndpoint(params *operations.GetDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointOK, error) {
	response, err := a.httpClient.GetDriverEndpoint(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverEndpoints(params *operations.GetDriverEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverEndpointsOK, error) {
	response, err := a.httpClient.GetDriverEndpoints(params, authInfo)
	return response, err
}
func (a *UsbClient) GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error) {
	response, err := a.httpClient.GetInfo(params, authInfo)
	return response, err
}
func (a *UsbClient) PingEndpoint(params *operations.PingDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverEndpointOK, error) {
	response, err := a.httpClient.PingDriverEndpoint(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error) {
	response, err := a.httpClient.UpdateCatalog(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDriverEndpoint(params *operations.UpdateDriverEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverEndpointOK, error) {
	response, err := a.httpClient.UpdateDriverEndpoint(params, authInfo)
	return response, err
}
func (a *UsbClient) SetTransport(transport runtime.ClientTransport) {
	a.httpClient.SetTransport(transport)
}

//GetDriverInstanceByName returns a *models.DriverInstance if found, else nil
func (a *UsbClient) GetDriverEndpointByName(instanceName string, authInfo runtime.ClientAuthInfoWriter) (*models.DriverEndpoint, error) {
	ret, err := a.GetDriverEndpoints(&operations.GetDriverEndpointsParams{}, authInfo)
	if err != nil {
		return nil, err
	}
	for _, d := range ret.Payload {
		if *d.Name == instanceName {
			return d, nil
		}
	}

	return nil, nil
}

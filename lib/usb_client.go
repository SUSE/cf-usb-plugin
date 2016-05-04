package lib

import (
	"github.com/go-openapi/runtime"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	strfmt "github.com/go-openapi/strfmt"
)

type UsbClient struct {
	httpClient *operations.Client
}

func NewUsbClient(transport runtime.ClientTransport, format strfmt.Registry) UsbClientInterface {
	return &UsbClient{httpClient: operations.New(transport, format)}
}

func (a *UsbClient) CreateDial(params *operations.CreateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDialCreated, error) {
	response, err := a.httpClient.CreateDial(params, authInfo)
	return response, err
}
func (a *UsbClient) CreateInstance(params *operations.CreateInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateInstanceCreated, error) {
	response, err := a.httpClient.CreateInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDial(params *operations.DeleteDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDialNoContent, error) {
	response, err := a.httpClient.DeleteDial(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteInstance(params *operations.DeleteInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteInstanceNoContent, error) {
	response, err := a.httpClient.DeleteInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) GetAllDials(params *operations.GetAllDialsParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetAllDialsOK, error) {
	response, err := a.httpClient.GetAllDials(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDial(params *operations.GetDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDialOK, error) {
	response, err := a.httpClient.GetDial(params, authInfo)
	return response, err
}
func (a *UsbClient) GetInstance(params *operations.GetInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInstanceOK, error) {
	response, err := a.httpClient.GetInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) GetInstances(params *operations.GetInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInstancesOK, error) {
	response, err := a.httpClient.GetInstances(params, authInfo)
	return response, err
}
func (a *UsbClient) GetInfo(params *operations.GetInfoParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetInfoOK, error) {
	response, err := a.httpClient.GetInfo(params, authInfo)
	return response, err
}
func (a *UsbClient) GetService(params *operations.GetServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceOK, error) {
	response, err := a.httpClient.GetService(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error) {
	response, err := a.httpClient.GetServiceByInstanceID(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServicePlan(params *operations.GetServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlanOK, error) {
	response, err := a.httpClient.GetServicePlan(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServicePlans(params *operations.GetServicePlansParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetServicePlansOK, error) {
	response, err := a.httpClient.GetServicePlans(params, authInfo)
	return response, err
}
func (a *UsbClient) PingInstance(params *operations.PingInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingInstanceOK, error) {
	response, err := a.httpClient.PingInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateCatalog(params *operations.UpdateCatalogParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateCatalogOK, error) {
	response, err := a.httpClient.UpdateCatalog(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDial(params *operations.UpdateDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDialOK, error) {
	response, err := a.httpClient.UpdateDial(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateInstance(params *operations.UpdateInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateInstanceOK, error) {
	response, err := a.httpClient.UpdateInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateService(params *operations.UpdateServiceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServiceOK, error) {
	response, err := a.httpClient.UpdateService(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateServicePlanOK, error) {
	response, err := a.httpClient.UpdateServicePlan(params, authInfo)
	return response, err
}
func (a *UsbClient) SetTransport(transport runtime.ClientTransport) {
	a.httpClient.SetTransport(transport)
}

//GetDriverInstanceByName returns a *models.DriverInstance if found, else nil
func (a *UsbClient) GetInstanceByName(authHeader runtime.ClientAuthInfoWriter, driverInstanceName string) (*models.Instance, error) {
	ret, err := a.GetInstances(&operations.GetInstancesParams{}, authHeader)
	if err != nil {
		return nil, err
	}
	for _, d := range ret.Payload {
		if *d.Name == driverInstanceName {
			return d, nil
		}
	}

	return nil, nil
}

//GetServiceByDriverInstanceID returns a service by driver instance id
func (a *UsbClient) GetServiceByDriverInstanceID(authInfo runtime.ClientAuthInfoWriter, driverInstanceID string) (*models.Service, error) {
	response, err := a.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{InstanceID: driverInstanceID}, authInfo)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

//GetPlanByID returns an existing plan by id
func (a *UsbClient) GetPlanByID(authInfo runtime.ClientAuthInfoWriter, planID string) (*models.Plan, error) {
	response, err := a.GetServicePlan(&operations.GetServicePlanParams{PlanID: planID}, authInfo)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

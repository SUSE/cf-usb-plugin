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
func (a *UsbClient) CreateDriver(params *operations.CreateDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDriverCreated, error) {
	response, err := a.httpClient.CreateDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) CreateDriverInstance(params *operations.CreateDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.CreateDriverInstanceCreated, error) {
	response, err := a.httpClient.CreateDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDial(params *operations.DeleteDialParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDialNoContent, error) {
	response, err := a.httpClient.DeleteDial(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDriver(params *operations.DeleteDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDriverNoContent, error) {
	response, err := a.httpClient.DeleteDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDriverInstance(params *operations.DeleteDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.DeleteDriverInstanceNoContent, error) {
	response, err := a.httpClient.DeleteDriverInstance(params, authInfo)
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
func (a *UsbClient) GetDialSchema(params *operations.GetDialSchemaParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDialSchemaOK, error) {
	response, err := a.httpClient.GetDialSchema(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriver(params *operations.GetDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverOK, error) {
	response, err := a.httpClient.GetDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverInstance(params *operations.GetDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverInstanceOK, error) {
	response, err := a.httpClient.GetDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverInstances(params *operations.GetDriverInstancesParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverInstancesOK, error) {
	response, err := a.httpClient.GetDriverInstances(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverSchema(params *operations.GetDriverSchemaParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriverSchemaOK, error) {
	response, err := a.httpClient.GetDriverSchema(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDrivers(params *operations.GetDriversParams, authInfo runtime.ClientAuthInfoWriter) (*operations.GetDriversOK, error) {
	response, err := a.httpClient.GetDrivers(params, authInfo)
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
func (a *UsbClient) PingDriverInstance(params *operations.PingDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.PingDriverInstanceOK, error) {
	response, err := a.httpClient.PingDriverInstance(params, authInfo)
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
func (a *UsbClient) UpdateDriver(params *operations.UpdateDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverOK, error) {
	response, err := a.httpClient.UpdateDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDriverInstance(params *operations.UpdateDriverInstanceParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UpdateDriverInstanceOK, error) {
	response, err := a.httpClient.UpdateDriverInstance(params, authInfo)
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
func (a *UsbClient) UploadDriver(params *operations.UploadDriverParams, authInfo runtime.ClientAuthInfoWriter) (*operations.UploadDriverOK, error) {
	response, err := a.httpClient.UploadDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) SetTransport(transport runtime.ClientTransport) {
	a.httpClient.SetTransport(transport)
}

//GetDriverByName returns a *model.driver if found else nil
func (a *UsbClient) GetDriverByName(authInfo runtime.ClientAuthInfoWriter, driverName string) (*models.Driver, error) {
	ret, err := a.GetDrivers(&operations.GetDriversParams{}, authInfo)
	if err != nil {
		return nil, err
	}

	var targetDriver *models.Driver

	for _, d := range ret.Payload {
		if *d.Name == driverName {
			targetDriver = d
		}
	}

	return targetDriver, nil
}

//GetDriverInstanceByName returns a *models.DriverInstance if found, else nil
func (a *UsbClient) GetDriverInstanceByName(authHeader runtime.ClientAuthInfoWriter, driverInstanceName string) (*models.DriverInstance, error) {
	ret, err := a.GetDrivers(&operations.GetDriversParams{}, authHeader)
	if err != nil {
		return nil, err
	}
	for _, d := range ret.Payload {
		for _, i := range d.DriverInstances {
			di, err := a.GetDriverInstance(&operations.GetDriverInstanceParams{DriverInstanceID: i}, authHeader)
			if err != nil {
				return nil, err
			}
			if *di.Payload.Name == driverInstanceName {
				return di.Payload, nil
			}
		}
	}

	return nil, nil
}

//GetServiceByDriverInstanceID returns a service by driver instance id
func (a *UsbClient) GetServiceByDriverInstanceID(authInfo runtime.ClientAuthInfoWriter, driverInstanceID string) (*models.Service, error) {
	response, err := a.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: driverInstanceID}, authInfo)
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

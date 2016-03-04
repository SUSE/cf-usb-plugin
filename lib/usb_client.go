package lib

import (
	client "github.com/go-swagger/go-swagger/client"
	operations "github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/go-swagger/go-swagger/strfmt"
)

type UsbClient struct {
	httpClient *operations.Client
}

func NewUsbClient(transport client.Transport, format strfmt.Registry) UsbClientInterface {
	return &UsbClient{httpClient: operations.New(transport, format)}
}

func (a *UsbClient) CreateDial(params *operations.CreateDialParams, authInfo client.AuthInfoWriter) (*operations.CreateDialCreated, error) {
	response, err := a.httpClient.CreateDial(params, authInfo)
	return response, err
}
func (a *UsbClient) CreateDriver(params *operations.CreateDriverParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverCreated, error) {
	response, err := a.httpClient.CreateDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) CreateDriverInstance(params *operations.CreateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.CreateDriverInstanceCreated, error) {
	response, err := a.httpClient.CreateDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDial(params *operations.DeleteDialParams, authInfo client.AuthInfoWriter) (*operations.DeleteDialNoContent, error) {
	response, err := a.httpClient.DeleteDial(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDriver(params *operations.DeleteDriverParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverNoContent, error) {
	response, err := a.httpClient.DeleteDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) DeleteDriverInstance(params *operations.DeleteDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.DeleteDriverInstanceNoContent, error) {
	response, err := a.httpClient.DeleteDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) GetAllDials(params *operations.GetAllDialsParams, authInfo client.AuthInfoWriter) (*operations.GetAllDialsOK, error) {
	response, err := a.httpClient.GetAllDials(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDial(params *operations.GetDialParams, authInfo client.AuthInfoWriter) (*operations.GetDialOK, error) {
	response, err := a.httpClient.GetDial(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDialSchema(params *operations.GetDialSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDialSchemaOK, error) {
	response, err := a.httpClient.GetDialSchema(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriver(params *operations.GetDriverParams, authInfo client.AuthInfoWriter) (*operations.GetDriverOK, error) {
	response, err := a.httpClient.GetDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverInstance(params *operations.GetDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstanceOK, error) {
	response, err := a.httpClient.GetDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverInstances(params *operations.GetDriverInstancesParams, authInfo client.AuthInfoWriter) (*operations.GetDriverInstancesOK, error) {
	response, err := a.httpClient.GetDriverInstances(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDriverSchema(params *operations.GetDriverSchemaParams, authInfo client.AuthInfoWriter) (*operations.GetDriverSchemaOK, error) {
	response, err := a.httpClient.GetDriverSchema(params, authInfo)
	return response, err
}
func (a *UsbClient) GetDrivers(params *operations.GetDriversParams, authInfo client.AuthInfoWriter) (*operations.GetDriversOK, error) {
	response, err := a.httpClient.GetDrivers(params, authInfo)
	return response, err
}
func (a *UsbClient) GetInfo(params *operations.GetInfoParams, authInfo client.AuthInfoWriter) (*operations.GetInfoOK, error) {
	response, err := a.httpClient.GetInfo(params, authInfo)
	return response, err
}
func (a *UsbClient) GetService(params *operations.GetServiceParams, authInfo client.AuthInfoWriter) (*operations.GetServiceOK, error) {
	response, err := a.httpClient.GetService(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServiceByInstanceID(params *operations.GetServiceByInstanceIDParams, authInfo client.AuthInfoWriter) (*operations.GetServiceByInstanceIDOK, error) {
	response, err := a.httpClient.GetServiceByInstanceID(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServicePlan(params *operations.GetServicePlanParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlanOK, error) {
	response, err := a.httpClient.GetServicePlan(params, authInfo)
	return response, err
}
func (a *UsbClient) GetServicePlans(params *operations.GetServicePlansParams, authInfo client.AuthInfoWriter) (*operations.GetServicePlansOK, error) {
	response, err := a.httpClient.GetServicePlans(params, authInfo)
	return response, err
}
func (a *UsbClient) PingDriverInstance(params *operations.PingDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.PingDriverInstanceOK, error) {
	response, err := a.httpClient.PingDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateCatalog(params *operations.UpdateCatalogParams, authInfo client.AuthInfoWriter) (*operations.UpdateCatalogOK, error) {
	response, err := a.httpClient.UpdateCatalog(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDial(params *operations.UpdateDialParams, authInfo client.AuthInfoWriter) (*operations.UpdateDialOK, error) {
	response, err := a.httpClient.UpdateDial(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDriver(params *operations.UpdateDriverParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverOK, error) {
	response, err := a.httpClient.UpdateDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateDriverInstance(params *operations.UpdateDriverInstanceParams, authInfo client.AuthInfoWriter) (*operations.UpdateDriverInstanceOK, error) {
	response, err := a.httpClient.UpdateDriverInstance(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateService(params *operations.UpdateServiceParams, authInfo client.AuthInfoWriter) (*operations.UpdateServiceOK, error) {
	response, err := a.httpClient.UpdateService(params, authInfo)
	return response, err
}
func (a *UsbClient) UpdateServicePlan(params *operations.UpdateServicePlanParams, authInfo client.AuthInfoWriter) (*operations.UpdateServicePlanOK, error) {
	response, err := a.httpClient.UpdateServicePlan(params, authInfo)
	return response, err
}
func (a *UsbClient) UploadDriver(params *operations.UploadDriverParams, authInfo client.AuthInfoWriter) (*operations.UploadDriverOK, error) {
	response, err := a.httpClient.UploadDriver(params, authInfo)
	return response, err
}
func (a *UsbClient) SetTransport(transport client.Transport) {
	a.httpClient.SetTransport(transport)
}

//GetDriverByName returns a *model.driver if found else nil
func (a *UsbClient) GetDriverByName(authInfo client.AuthInfoWriter, driverName string) (*models.Driver, error) {
	ret, err := a.GetDrivers(&operations.GetDriversParams{}, authInfo)
	if err != nil {
		return nil, err
	}

	var targetDriver *models.Driver

	for _, d := range ret.Payload {
		if d.Name == driverName {
			targetDriver = d
		}
	}

	return targetDriver, nil
}

//GetDriverInstanceByName returns a *models.DriverInstance if found, else nil
func (a *UsbClient) GetDriverInstanceByName(authHeader client.AuthInfoWriter, driverInstanceName string) (*models.DriverInstance, error) {
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
			if di.Payload.Name == driverInstanceName {
				return di.Payload, nil
			}
		}
	}

	return nil, nil
}

//GetServiceByDriverInstanceID returns a service by driver instance id
func (a *UsbClient) GetServiceByDriverInstanceID(authInfo client.AuthInfoWriter, driverInstanceID string) (*models.Service, error) {
	response, err := a.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: driverInstanceID}, authInfo)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

//GetPlanByID returns an existing plan by id
func (a *UsbClient) GetPlanByID(authInfo client.AuthInfoWriter, planID string) (*models.Plan, error) {
	response, err := a.GetServicePlan(&operations.GetServicePlanParams{PlanID: planID}, authInfo)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

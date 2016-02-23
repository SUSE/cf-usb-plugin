package commands

import (
	"fmt"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//ServiceInterface exposes service commands
type ServiceInterface interface {
	GetServiceByDriverInstanceID(swaggerclient.AuthInfoWriter, string) (*models.Service, error)
	Update(swaggerclient.AuthInfoWriter, *models.Service) (string, error)
}

//ServiceCommands struct
type ServiceCommands struct {
	instanceCommands InstanceInterface
	httpClient       lib.UsbClientInterface
}

//NewServiceCommands returns a ServiceCommands object
func NewServiceCommands(httpClient lib.UsbClientInterface, instance InstanceInterface) ServiceInterface {
	return &ServiceCommands{httpClient: httpClient, instanceCommands: instance}
}

//GetServiceByID returns a service by driver instance id
func (c *ServiceCommands) GetServiceByDriverInstanceID(bearer swaggerclient.AuthInfoWriter, driverInstanceID string) (*models.Service, error) {
	response, err := c.httpClient.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: driverInstanceID}, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

//Update - updates a service's details

func (c *ServiceCommands) Update(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	instance := c.instanceCommands.GetDriverInstanceByName(bearer, args[0])
	if instance == nil {
		fmt.Println("Driver instance not found")
		return "", nil
	}

	params := operations.NewUpdateServiceParams()
	params.ServiceID = *service.ID
	params.Service = service

	fmt.Println("bindable:", *service.Bindable)
	fmt.Println("bindable:", service.Bindable)

	response, err := c.httpClient.UpdateService(params, bearer)
	if err != nil {
		return "", err
	}

	return *response.Payload.ID, nil
}

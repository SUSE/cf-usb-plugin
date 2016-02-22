package commands

import (
	"fmt"
	"strings"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//ServiceInterface exposes service commands
type ServiceInterface interface {
	GetServiceByID(swaggerclient.AuthInfoWriter, string) (*models.Service, error)
	Update(swaggerclient.AuthInfoWriter, []string) (string, error)
}

//ServiceCommands struct
type ServiceCommands struct {
	httpClient *operations.Client
}

//NewServiceCommands returns a ServiceCommands object
func NewServiceCommands(httpClient *operations.Client) ServiceInterface {
	return &ServiceCommands{httpClient: httpClient}
}

//GetServiceByID returns a service by driver instance id
func (c *ServiceCommands) GetServiceByID(bearer swaggerclient.AuthInfoWriter, driverID string) (*models.Service, error) {
	response, err := c.httpClient.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: driverID}, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

//Update - updates a service's details
func (c *ServiceCommands) Update(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	instance := getDriverInstanceByName(c.httpClient, bearer, args[0])
	if instance == nil {
		fmt.Println("Driver instance not found")
		return "", nil
	}

	params := operations.NewUpdateServiceParams()
	params.ServiceID = *instance.Service

	var service models.Service
	service.DriverInstanceID = *instance.ID

	bindable := true
	if strings.ToLower(strings.Trim(args[1], " ")) == "n" {
		bindable = false
	}

	service.Bindable = &bindable
	service.Name = args[2]
	service.Description = &args[3]
	service.Tags = strings.Split(args[4], ",")

	params.Service = &service

	response, err := c.httpClient.UpdateService(params, bearer)
	if err != nil {
		fmt.Println("ERROR:", err)
		return "", err
	}

	return *response.Payload.ID, nil
}

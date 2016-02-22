package commands

import (
	"fmt"
	"strings"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type ServiceInterface interface {
	GetServiceById(swaggerclient.AuthInfoWriter, string) (*models.Service, error)
	Update(swaggerclient.AuthInfoWriter, []string) (string, error)
}

type ServiceCommands struct {
	httpClient *operations.Client
}

func NewServiceCommands(httpClient *operations.Client) ServiceInterface {
	return &ServiceCommands{httpClient: httpClient}
}

func (c *ServiceCommands) GetServiceById(bearer swaggerclient.AuthInfoWriter, driverId string) (*models.Service, error) {
	response, err := c.httpClient.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: driverId}, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

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

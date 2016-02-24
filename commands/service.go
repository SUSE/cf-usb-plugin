package commands

import (
	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//ServiceInterface exposes service commands
type ServiceInterface interface {
	Update(swaggerclient.AuthInfoWriter, *models.Service) (string, error)
}

//ServiceCommands struct
type ServiceCommands struct {
	httpClient lib.UsbClientInterface
}

//NewServiceCommands returns a ServiceCommands object
func NewServiceCommands(httpClient lib.UsbClientInterface) ServiceInterface {
	return &ServiceCommands{httpClient: httpClient}
}

//Update - updates a service's details
func (c *ServiceCommands) Update(bearer swaggerclient.AuthInfoWriter, service *models.Service) (string, error) {
	params := operations.NewUpdateServiceParams()
	params.ServiceID = *service.ID
	params.Service = service

	response, err := c.httpClient.UpdateService(params, bearer)
	if err != nil {
		return "", err
	}

	return *response.Payload.ID, nil
}

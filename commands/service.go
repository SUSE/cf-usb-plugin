package commands

import (
	"github.com/go-openapi/runtime"

	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//ServiceInterface exposes service commands
type ServiceInterface interface {
	Update(runtime.ClientAuthInfoWriter, *models.Service) (string, error)
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
func (c *ServiceCommands) Update(bearer runtime.ClientAuthInfoWriter, service *models.Service) (string, error) {
	params := operations.NewUpdateServiceParams()
	params.ServiceID = service.ID
	params.Service = service

	response, err := c.httpClient.UpdateService(params, bearer)
	if err != nil {
		return "", err
	}

	return response.Payload.ID, nil
}

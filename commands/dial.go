package commands

import (
	"fmt"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//DialInterface exposes list dials command
type DialInterface interface {
	List(swaggerclient.AuthInfoWriter, string) ([]*models.Dial, error)
}

//DialCommands struct
type DialCommands struct {
	httpClient lib.UsbClientInterface
}

//NewDialCommands returns a DialCommands object
func NewDialCommands(httpClient lib.UsbClientInterface) DialInterface {
	return &DialCommands{httpClient: httpClient}
}

//List dials of an instance
func (c *DialCommands) List(bearer swaggerclient.AuthInfoWriter, instanceName string) ([]*models.Dial, error) {
	instance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return nil, err
	}
	if instance == nil {
		fmt.Println("Driver instance not found")
		return nil, nil
	}

	params := operations.NewGetAllDialsParams()
	params.DriverInstanceID = instance.ID

	response, err := c.httpClient.GetAllDials(params, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

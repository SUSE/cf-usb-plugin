package commands

import (
	"fmt"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type DialInterface interface {
	List(swaggerclient.AuthInfoWriter, string) ([]*models.Dial, error)
}

type DialCommands struct {
	httpClient *operations.Client
}

func NewDialCommands(httpClient *operations.Client) DialInterface {
	return &DialCommands{httpClient: httpClient}
}

func (c *DialCommands) List(bearer swaggerclient.AuthInfoWriter, instanceName string) ([]*models.Dial, error) {
	instance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
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

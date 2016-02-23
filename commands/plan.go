package commands

import (
	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//PlanInterface exposes get plan by id command
type PlanInterface interface {
	GetPlanByID(swaggerclient.AuthInfoWriter, string) (*models.Plan, error)
}

//PlanCommands struct
type PlanCommands struct {
	httpClient lib.UsbClientInterface
}

//NewPlanCommands returns a PlanCommands object
func NewPlanCommands(httpClient lib.UsbClientInterface) PlanInterface {
	return &PlanCommands{httpClient: httpClient}
}

//GetPlanByID returns an existing plan by id
func (c *PlanCommands) GetPlanByID(bearer swaggerclient.AuthInfoWriter, planID string) (*models.Plan, error) {
	response, err := c.httpClient.GetServicePlan(&operations.GetServicePlanParams{PlanID: planID}, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

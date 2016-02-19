package commands

import (
	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type PlanInterface interface {
	GetPlanById(swaggerclient.AuthInfoWriter, string) (*models.Plan, error)
}

type PlanCommands struct {
	httpClient *operations.Client
}

func NewPlanCommands(httpClient *operations.Client) PlanInterface {
	return &PlanCommands{httpClient: httpClient}
}

func (c *PlanCommands) GetPlanById(bearer swaggerclient.AuthInfoWriter, planId string) (*models.Plan, error) {
	response, err := c.httpClient.GetServicePlan(&operations.GetServicePlanParams{PlanID: planId}, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

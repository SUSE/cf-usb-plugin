package commands

import (
	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

type InfoInterface interface {
	GetInfo(swaggerclient.AuthInfoWriter) (*models.Info, error)
}

type InfoCommands struct {
	httpClient *operations.Client
}

func NewInfoCommands(httpClient *operations.Client) InfoInterface {
	return &InfoCommands{httpClient: httpClient}
}

func (c *InfoCommands) GetInfo(bearer swaggerclient.AuthInfoWriter) (*models.Info, error) {
	infoResp, err := c.httpClient.GetInfo(operations.NewGetInfoParams(), bearer)
	if err != nil {
		return nil, err
	}

	return infoResp.Payload, nil
}

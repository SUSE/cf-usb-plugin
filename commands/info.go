package commands

import (
	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//InfoInterface exposes GetInfo command
type InfoInterface interface {
	GetInfo(swaggerclient.AuthInfoWriter) (*models.Info, error)
}

//InfoCommands struct
type InfoCommands struct {
	httpClient lib.UsbClientInterface
}

//NewInfoCommands returns an InfoCommands object
func NewInfoCommands(httpClient lib.UsbClientInterface) InfoInterface {
	return &InfoCommands{httpClient: httpClient}
}

//GetInfo - retruns usb information
func (c *InfoCommands) GetInfo(bearer swaggerclient.AuthInfoWriter) (*models.Info, error) {
	infoResp, err := c.httpClient.GetInfo(operations.NewGetInfoParams(), bearer)
	if err != nil {
		return nil, err
	}

	return infoResp.Payload, nil
}

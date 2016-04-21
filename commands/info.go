package commands

import (
	"github.com/go-openapi/runtime"
	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//InfoInterface exposes GetInfo command
type InfoInterface interface {
	GetInfo(runtime.ClientAuthInfoWriter) (*models.Info, error)
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
func (c *InfoCommands) GetInfo(bearer runtime.ClientAuthInfoWriter) (*models.Info, error) {
	infoResp, err := c.httpClient.GetInfo(operations.NewGetInfoParams(), bearer)
	if err != nil {
		return nil, err
	}

	return infoResp.Payload, nil
}

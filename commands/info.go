package commands

import (
	"github.com/go-openapi/runtime"
	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	httptransport "github.com/go-openapi/runtime/client"
)

//InfoInterface exposes GetInfo command
type InfoInterface interface {
	GetInfo() (*models.Info, error)
}

//InfoCommands struct
type InfoCommands struct {
	httpClient lib.UsbClientInterface
	token      runtime.ClientAuthInfoWriter
}

//NewInfoCommands returns an InfoCommands object
func NewInfoCommands(httpClient lib.UsbClientInterface, bearer string) InfoInterface {
	return &InfoCommands{httpClient: httpClient, token: httptransport.BearerToken(bearer)}
}

//GetInfo - retruns usb information
func (c *InfoCommands) GetInfo() (*models.Info, error) {
	infoResp, err := c.httpClient.GetInfo(operations.NewGetInfoParams(), c.token)
	if err != nil {
		return nil, err
	}

	return infoResp.Payload, nil
}

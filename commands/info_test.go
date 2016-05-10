package commands_test

import (
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	"testing"

	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/stretchr/testify/assert"
)

func Test_GetInfo(t *testing.T) {
	assert := assert.New(t)

	var infoResponse operations.GetInfoOK
	var info models.Info
	apiVersion := "testAPI"
	usbVersion := "testUSB"
	info.BrokerAPIVersion = &apiVersion
	info.UsbVersion = &usbVersion
	infoResponse.Payload = &info

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	usbClientMock.GetInfoReturns(&infoResponse, nil)

	bearer := httptransport.BearerToken("testToken")

	infoCommands := commands.NewInfoCommands(usbClientMock)

	response, err := infoCommands.GetInfo(bearer)

	assert.NotNil(response)
	assert.Equal("testAPI", *response.BrokerAPIVersion)
	assert.Equal("testUSB", *response.UsbVersion)
	assert.NoError(err)
}

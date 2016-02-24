package commands_test

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetInfo(t *testing.T) {
	assert := assert.New(t)

	var infoResponse operations.GetInfoOK
	var info models.Info
	info.BrokerAPIVersion = "testAPI"
	info.UsbVersion = "testUSB"
	infoResponse.Payload = &info

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	usbClientMock.GetInfoReturns(&infoResponse, nil)

	bearer := httptransport.BearerToken("testToken")

	infoCommands := commands.NewInfoCommands(usbClientMock)

	response, err := infoCommands.GetInfo(bearer)

	assert.NotNil(response)
	assert.Equal(response.BrokerAPIVersion, "testAPI")
	assert.Equal(response.UsbVersion, "testUSB")
	assert.NoError(err)
}

package commands_test

import (
	"github.com/SUSE/cf-usb-plugin/lib/client/operations"
	"github.com/SUSE/cf-usb-plugin/lib/models"

	"github.com/SUSE/cf-usb-plugin/commands"

	"testing"

	fakeUsbClient "github.com/SUSE/cf-usb-plugin/lib/fakes"
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

	infoCommands := commands.NewInfoCommands(usbClientMock, testBearer)

	response, err := infoCommands.GetInfo()

	assert.NotNil(response)
	assert.Equal("testAPI", *response.BrokerAPIVersion)
	assert.Equal("testUSB", *response.UsbVersion)
	assert.NoError(err)
}

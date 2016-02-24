package commands_test

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	"testing"

	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateService(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	bearer := httptransport.BearerToken("testToken")

	var service models.Service
	service.Name = "test"
	serviceID := "testID"
	service.ID = &serviceID

	var result operations.UpdateServiceOK
	result.Payload = &service

	usbClientMock.UpdateServiceReturns(&result, nil)

	serviceCommands := commands.NewServiceCommands(usbClientMock)

	response, err := serviceCommands.Update(bearer, &service)
	t.Log(response)
	assert.NoError(err)
}

package commands_test

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	fakeCommands "github.com/hpcloud/cf-plugin-usb/commands/fakes"
	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ListDials(t *testing.T) {
	assert := assert.New(t)

	var getAllDialsResponseOK operations.GetAllDialsOK

	var getAllDialsResponse []*models.Dial
	var fake models.Dial
	fakeID := "fake"
	fakePlanID := "fakePlanID"
	fake.ID = &fakeID
	fake.DriverInstanceID = "testInstanceID"
	fake.Plan = &fakePlanID
	fake.Configuration = map[string]string{"test": "test"}
	getAllDialsResponse = append(getAllDialsResponse, &fake)

	getAllDialsResponseOK.Payload = getAllDialsResponse

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)
	fakeInstanceCommands := new(fakeCommands.FakeInstanceInterface)

	usbClientMock.GetAllDialsReturns(&getAllDialsResponseOK, nil)

	var instance models.DriverInstance
	instanceID := "testInstanceID"
	instance.ID = &instanceID
	instance.Name = "testInstance"
	instance.Dials = append(instance.Dials, "fake")

	bearer := httptransport.BearerToken("testToken")

	fakeInstanceCommands.GetDriverInstanceByNameReturns(&instance)

	dialCommands := commands.NewDialCommands(usbClientMock, fakeInstanceCommands)

	response, err := dialCommands.List(bearer, "testInstance")
	for _, d := range response {
		t.Log(*d.ID)
	}
	assert.NotNil(response)
	assert.NoError(err)
}

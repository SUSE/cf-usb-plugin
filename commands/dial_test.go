package commands

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	fakes "github.com/hpcloud/cf-plugin-usb/lib/fakes"
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

	usbClientMock := new(fakes.FakeUsbClientInterface)

	usbClientMock.GetAllDialsReturns(&getAllDialsResponseOK, nil)

	var instance models.DriverInstance
	instanceID := "testInstanceID"
	instance.ID = &instanceID
	instance.Name = "testInstance"
	instance.Dials = append(instance.Dials, "fake")

	usbClientMock.GetDriverInstanceByNameReturns(&instance, nil)
	bearer := httptransport.BearerToken("testToken")

	dialCommands := NewDialCommands(usbClientMock)

	response, err := dialCommands.List(bearer, "testInstance")
	for _, d := range response {
		t.Log(*d.ID)
	}
	assert.NotNil(response)
	assert.NoError(err)
}

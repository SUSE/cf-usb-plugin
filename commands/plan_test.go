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

func Test_GetPlanByID(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	bearer := httptransport.BearerToken("testToken")
	var planResult operations.GetServicePlanOK
	var plan models.Plan
	planID := "testPlanID"
	plan.ID = &planID
	plan.Name = "testPlanName"
	planResult.Payload = &plan
	usbClientMock.GetServicePlanReturns(&planResult, nil)

	planCommands := commands.NewPlanCommands(usbClientMock)

	response, err := planCommands.GetPlanByID(bearer, "testPlanID")
	assert.NotNil(response)
	assert.NoError(err)
}

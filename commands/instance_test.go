package commands_test

import (
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	"encoding/json"
	"testing"

	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/stretchr/testify/assert"
)

var testBearer = "testToken"

func Test_CreateDriverInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	instanceCommands := commands.NewInstanceCommands(usbClientMock, testBearer)

	testID := "testID"

	var createdInstance models.DriverEndpoint
	var createResult operations.RegisterDriverEndpointCreated

	name := "testInstance"
	id := "testID"
	createdInstance.Name = &name
	createdInstance.ID = id
	createResult.Payload = &createdInstance
	createdInstance.ID = id
	usbClientMock.RegisterDriverEndpointReturns(&createResult, nil)

	var metadata *json.RawMessage
	meta := json.RawMessage(`{"display_name":"name"}`)
	metadata = &meta

	response, err := instanceCommands.Create("testDriver", "http://127.0.0.1", "key", metadata)
	assert.Equal(response, testID)
	assert.NoError(err)
}

func Test_DeleteInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	instanceCommands := commands.NewInstanceCommands(usbClientMock, testBearer)

	testID := "testID"
	name := "testDriver"

	var testDriverInstance models.DriverEndpoint
	testDriverInstance.Name = &name
	testDriverInstance.ID = testID

	usbClientMock.GetDriverEndpointByNameReturns(&testDriverInstance, nil)

	var deleteResult operations.UnregisterDriverInstanceNoContent

	usbClientMock.UnregisterDriverEndpointReturns(&deleteResult, nil)

	_, err := instanceCommands.Delete("testID")
	assert.NoError(err)
}

func Test_UpdateInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	instanceCommands := commands.NewInstanceCommands(usbClientMock, testBearer)

	testID := "testID"
	testName := "testDriver"

	var testDriverResult operations.GetDriverEndpointOK

	var testDriver models.DriverEndpoint
	testDriver.Name = &testName
	testDriver.ID = testID
	testDriver.AuthenticationKey = "auth"

	testDriverResult.Payload = &testDriver

	usbClientMock.GetDriverEndpointReturns(&testDriverResult, nil)

	var oldInstance models.DriverEndpoint
	oldInstance.Name = &testName
	oldInstance.ID = testID

	var updateResult operations.UpdateDriverEndpointOK
	var upInstance models.DriverEndpoint
	newName := "testInstanceUpdate"

	upInstance.Name = &newName
	upInstance.ID = "testID"

	updateResult.Payload = &upInstance

	usbClientMock.UpdateDriverEndpointReturns(&updateResult, nil)

	usbClientMock.GetDriverEndpointByNameReturns(&oldInstance, nil)

	metadata := json.RawMessage(`{"display_name":"name"}`)

	response, err := instanceCommands.Update("testDriver", &metadata)
	assert.NotEqual(response, oldInstance.Name)
	assert.NoError(err)
}

func Test_ListDriverInstances(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	instanceCommands := commands.NewInstanceCommands(usbClientMock, testBearer)

	var result operations.GetDriverEndpointsOK

	var instances []*models.DriverEndpoint

	var instance models.DriverEndpoint

	name := "testInstance"

	instance.Name = &name
	instance.ID = "testID"

	instances = append(instances, &instance)

	result.Payload = instances
	usbClientMock.GetDriverEndpointsReturns(&result, nil)

	response, err := instanceCommands.List()

	assert.NotNil(response)
	assert.NoError(err)
}

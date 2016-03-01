package commands_test

import (
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	"github.com/hpcloud/cf-plugin-usb/commands"

	testterm "github.com/cloudfoundry/cli/testhelpers/terminal"
	fakeUsbClient "github.com/hpcloud/cf-plugin-usb/lib/fakes"
	"github.com/hpcloud/cf-plugin-usb/lib/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateDriverInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)
	ui := new(testterm.FakeUI)

	schemaParser := schema.NewSchemaParser(ui)
	instanceCommands := commands.NewInstanceCommands(usbClientMock, schemaParser)

	bearer := httptransport.BearerToken("testToken")
	testID := "testID"

	var testDriver models.Driver
	testDriver.Name = "testDriver"
	testDriver.ID = &testID
	testDriver.DriverType = "testType"

	usbClientMock.GetDriverByNameReturns(&testDriver, nil)

	var createResult operations.CreateDriverInstanceCreated
	var createdInstance models.DriverInstance
	createdInstance.Name = "testInstance"
	createdInstance.DriverID = "testID"
	createResult.Payload = &createdInstance
	createdInstance.ID = &testID
	usbClientMock.CreateDriverInstanceReturns(&createResult, nil)

	response, err := instanceCommands.Create(bearer, []string{"testDriver", "testInstance", "configValue", `{"a":"b"}`})
	assert.Equal(response, testID)
	assert.NoError(err)
}

func Test_DeleteInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)
	ui := new(testterm.FakeUI)

	schemaParser := schema.NewSchemaParser(ui)
	instanceCommands := commands.NewInstanceCommands(usbClientMock, schemaParser)

	bearer := httptransport.BearerToken("testToken")
	testID := "testID"

	var testDriver models.Driver
	testDriver.Name = "testDriver"
	testDriver.ID = &testID
	testDriver.DriverType = "testType"

	usbClientMock.GetDriverByNameReturns(&testDriver, nil)

	var deleteResult operations.DeleteDriverInstanceNoContent

	usbClientMock.DeleteDriverInstanceReturns(&deleteResult, nil)

	_, err := instanceCommands.Delete(bearer, "testID")
	assert.NoError(err)
}

func Test_UpdateInstance(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)
	ui := new(testterm.FakeUI)

	schemaParser := schema.NewSchemaParser(ui)
	instanceCommands := commands.NewInstanceCommands(usbClientMock, schemaParser)

	bearer := httptransport.BearerToken("testToken")
	testID := "testID"

	var testDriverResult operations.GetDriverOK

	var testDriver models.Driver
	testDriver.Name = "testDriver"
	testDriver.ID = &testID
	testDriver.DriverType = "testType"

	testDriverResult.Payload = &testDriver

	usbClientMock.GetDriverReturns(&testDriverResult, nil)

	var oldInstance models.DriverInstance
	oldInstance.Name = "testInstance"
	oldInstance.DriverID = "testID"
	oldInstance.ID = &testID

	var updateResult operations.UpdateDriverInstanceOK
	var upInstance models.DriverInstance
	upInstance.Name = "testInstanceUpdate"
	upInstance.DriverID = "testID"
	upInstance.ID = &testID

	updateResult.Payload = &upInstance

	usbClientMock.UpdateDriverInstanceReturns(&updateResult, nil)

	usbClientMock.GetDriverInstanceByNameReturns(&oldInstance, nil)

	response, err := instanceCommands.Update(bearer, []string{"testDriver", "testInstanceUpdate", "configValue", `{"a":"b"}`})
	assert.NotEqual(response, oldInstance.Name)
	assert.NoError(err)
}

func Test_ListDriverInstances(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)
	ui := new(testterm.FakeUI)

	schemaParser := schema.NewSchemaParser(ui)
	instanceCommands := commands.NewInstanceCommands(usbClientMock, schemaParser)

	bearer := httptransport.BearerToken("testToken")
	testID := "testID"

	var testDriver models.Driver
	testDriver.Name = "testDriver"
	testDriver.ID = &testID
	testDriver.DriverType = "testType"

	usbClientMock.GetDriverByNameReturns(&testDriver, nil)

	var instancesResult operations.GetDriverInstancesOK

	var instaces []*models.DriverInstance

	var instance models.DriverInstance
	instance.Name = "testInstance"
	instance.DriverID = "testID"
	testInstanceID := "testInstanceID"
	instance.ID = &testInstanceID

	instaces = append(instaces, &instance)

	instancesResult.Payload = instaces
	usbClientMock.GetDriverInstancesReturns(&instancesResult, nil)

	response, err := instanceCommands.List(bearer, "testID")

	assert.NotNil(response)
	assert.NoError(err)
}

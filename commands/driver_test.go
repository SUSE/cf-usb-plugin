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

func Test_CreateDriver(t *testing.T) {
	assert := assert.New(t)
	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	var driverResonse operations.CreateDriverCreated

	var driver models.Driver
	driver.DriverType = "testType"
	driver.Name = "testName"
	driverID := "testID"
	driver.ID = &driverID
	driver.DriverInstances = []string{"testInstanceID"}

	driverResonse.Payload = &driver
	usbClientMock.CreateDriverReturns(&driverResonse, nil)

	driverCommands := commands.NewDriverCommands(usbClientMock)

	bearer := httptransport.BearerToken("testToken")

	result, err := driverCommands.Create(bearer, []string{"testType", "testName"})
	assert.Equal("testID", result)
	assert.NoError(err)
}

func Test_DeleteDriver(t *testing.T) {
	assert := assert.New(t)

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	var driverResonse operations.GetDriversOK

	var response []*models.Driver

	var driver models.Driver
	driver.DriverType = "testType"
	driver.Name = "test"
	driverID := "testID"
	driver.ID = &driverID
	driver.DriverInstances = []string{"testInstanceID"}

	response = append(response, &driver)

	driverResonse.Payload = response
	usbClientMock.GetDriversReturns(&driverResonse, nil)

	var deleteResponse operations.DeleteDriverNoContent

	usbClientMock.GetDriverByNameReturns(&driver, nil)
	usbClientMock.DeleteDriverReturns(&deleteResponse, nil)

	driverCommands := commands.NewDriverCommands(usbClientMock)

	bearer := httptransport.BearerToken("testToken")

	result, err := driverCommands.Delete(bearer, "test")
	assert.Equal("testID", result)
	assert.NoError(err)
}

func Test_UpdateDriver(t *testing.T) {
	assert := assert.New(t)

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	var driverResonse operations.GetDriversOK

	var response []*models.Driver

	var driver models.Driver
	driver.DriverType = "testType"
	driver.Name = "test"
	driverID := "testID"
	driver.ID = &driverID
	driver.DriverInstances = []string{"testInstanceID"}

	response = append(response, &driver)

	driverResonse.Payload = response

	usbClientMock.GetDriverByNameReturns(&driver, nil)
	usbClientMock.GetDriversReturns(&driverResonse, nil)

	var updateResponse operations.UpdateDriverOK

	var updated models.Driver
	updated = driver
	updated.Name = "testUpdated"

	updateResponse.Payload = &updated

	usbClientMock.UpdateDriverReturns(&updateResponse, nil)

	driverCommands := commands.NewDriverCommands(usbClientMock)

	bearer := httptransport.BearerToken("testToken")

	result, err := driverCommands.Update(bearer, []string{"test", "testUpdated"})
	assert.Equal("testUpdated", result)
	assert.NoError(err)
}

func Test_ListDrivers(t *testing.T) {
	assert := assert.New(t)

	usbClientMock := new(fakeUsbClient.FakeUsbClientInterface)

	var driverResonse operations.GetDriversOK

	var response []*models.Driver

	var driver models.Driver
	driver.DriverType = "testType"
	driver.Name = "test"
	driverID := "testID"
	driver.ID = &driverID
	driver.DriverInstances = []string{"testInstanceID"}

	response = append(response, &driver)

	driverResonse.Payload = response
	usbClientMock.GetDriversReturns(&driverResonse, nil)

	driverCommands := commands.NewDriverCommands(usbClientMock)

	bearer := httptransport.BearerToken("testToken")

	result, err := driverCommands.List(bearer)
	for _, d := range result {
		t.Log(d.Name)
	}
	assert.NoError(err)
}

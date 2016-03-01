package commands

import (
	"os"
	"sort"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//DriverInterface exposes driver commands
type DriverInterface interface {
	Create(swaggerclient.AuthInfoWriter, []string) (string, error)
	Delete(swaggerclient.AuthInfoWriter, string) (string, error)
	Update(swaggerclient.AuthInfoWriter, []string) (string, error)
	List(swaggerclient.AuthInfoWriter) ([]*models.Driver, error)
}

//DriverCommands struct
type DriverCommands struct {
	httpClient lib.UsbClientInterface
}

//NewDriverCommands returns a DriverCommands object
func NewDriverCommands(httpClient lib.UsbClientInterface) DriverInterface {
	return &DriverCommands{httpClient: httpClient}
}

//Create - creates a new driver
func (c *DriverCommands) Create(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	// if bits path specified, check if exists
	if len(args) == 3 {
		if _, err := os.Stat(args[2]); err != nil {
			return "", err
		}
	}

	var driver models.Driver
	driver.DriverType = args[0]
	driver.Name = args[1]

	params := operations.NewCreateDriverParams()
	params.Driver = &driver

	response, err := c.httpClient.CreateDriver(params, bearer)
	if err != nil {
		return "", err
	}

	filePath := ""
	if len(args) == 3 {
		filePath = args[2]

		sha, err := getFileSha(filePath)
		if err != nil {
			return "", err
		}

		file, err := os.Open(filePath)
		if err != nil {
			return "", err
		}

		var uploadParams operations.UploadDriverParams

		uploadParams.DriverID = *response.Payload.ID
		uploadParams.File = *file
		uploadParams.Sha = sha

		_, err = c.httpClient.UploadDriver(&uploadParams, bearer)
		if err != nil {
			return "", err
		}
	}

	return *response.Payload.ID, nil
}

//Delete - deletes an existing driver
func (c *DriverCommands) Delete(bearer swaggerclient.AuthInfoWriter, driverName string) (string, error) {

	driver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if err != nil {
		return "", err
	}
	if driver == nil {
		return "", nil
	}

	params := operations.NewDeleteDriverParams()
	params.DriverID = *driver.ID

	_, err = c.httpClient.DeleteDriver(params, bearer)
	if err != nil {
		return "", err
	}

	return *driver.ID, nil
}

//Update - updates an existing driver
func (c *DriverCommands) Update(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	oldName := args[0]
	newName := args[1]

	driver, err := c.httpClient.GetDriverByName(bearer, oldName)
	if err != nil {
		return "", err
	}
	if driver == nil {
		return "", nil
	}
	driver.Name = newName

	params := operations.NewUpdateDriverParams()
	params.DriverID = *driver.ID
	params.Driver = driver

	response, err := c.httpClient.UpdateDriver(params, bearer)
	if err != nil {
		return "", err
	}

	return response.Payload.Name, nil
}

type driverSorter []*models.Driver

func (a driverSorter) Len() int           { return len(a) }
func (a driverSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a driverSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

//List - lists existing drivers
func (c *DriverCommands) List(bearer swaggerclient.AuthInfoWriter) ([]*models.Driver, error) {
	response, err := c.httpClient.GetDrivers(operations.NewGetDriversParams(), bearer)
	if err != nil {
		return nil, err
	}
	sort.Sort(driverSorter(response.Payload))
	return response.Payload, nil
}

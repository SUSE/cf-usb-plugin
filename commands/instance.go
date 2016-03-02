package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
	"github.com/hpcloud/cf-plugin-usb/lib/schema"
)

//InstanceInterface exposes instances commands
type InstanceInterface interface {
	Create(swaggerclient.AuthInfoWriter, []string) (string, error)
	Delete(swaggerclient.AuthInfoWriter, string) (string, error)
	Update(swaggerclient.AuthInfoWriter, []string) (string, error)
	List(swaggerclient.AuthInfoWriter, string) ([]*models.DriverInstance, error)
}

//InstanceCommands struct
type InstanceCommands struct {
	httpClient   lib.UsbClientInterface
	schemaParser *schema.SchemaParser
}

//NewInstanceCommands - returns an InstanceCommands object
func NewInstanceCommands(httpClient lib.UsbClientInterface, schemaParser *schema.SchemaParser) InstanceInterface {
	return &InstanceCommands{httpClient: httpClient, schemaParser: schemaParser}
}

//Create - creates a new driver instance
func (c *InstanceCommands) Create(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	driverName := args[0]
	instanceName := args[1]

	targetDriver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if targetDriver == nil {
		fmt.Println("Driver not found")
		return "", nil
	}

	var driverConfig map[string]interface{}

	if len(args) == 4 {
		method := args[2]
		configValue := args[3]

		if method == "jsonfile" {
			fileContent, err := ioutil.ReadFile(configValue)
			if err != nil {
				return "", fmt.Errorf("Unable to read configuration file. %s", err.Error())
			}
			configValue = string(fileContent)
		}

		if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
			return "", fmt.Errorf("Invalid JSON format %s", err.Error())
		}

	} else if len(args) == 2 {
		configSchema, err := c.httpClient.GetDriverSchema(&operations.GetDriverSchemaParams{DriverID: *targetDriver.ID}, bearer)
		if err != nil {
			return "", err
		}

		configValue, err := c.schemaParser.ParseSchema(string(configSchema.Payload))
		if err != nil {
			return "", err
		}

		if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
			return "", fmt.Errorf("Invalid JSON format %s", err.Error())
		}
	}

	newDriver := models.DriverInstance{
		Name:          instanceName,
		DriverID:      *targetDriver.ID,
		Configuration: driverConfig,
	}

	response, err := c.httpClient.CreateDriverInstance(&operations.CreateDriverInstanceParams{DriverInstance: &newDriver}, bearer)
	if err != nil {
		return "", err
	}

	return *response.Payload.ID, nil
}

//Delete - deletes an existing driver instance
func (c *InstanceCommands) Delete(bearer swaggerclient.AuthInfoWriter, instanceName string) (string, error) {
	instance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return "", err
	}
	if instance == nil {
		return "", nil
	}

	params := operations.NewDeleteDriverInstanceParams()
	params.DriverInstanceID = *instance.ID

	_, err = c.httpClient.DeleteDriverInstance(params, bearer)
	if err != nil {
		return "", err
	}

	return *instance.ID, nil
}

//Update - updates an existing driver instance
func (c *InstanceCommands) Update(bearer swaggerclient.AuthInfoWriter, args []string) (string, error) {
	driverName := args[0]
	instanceName := args[1]

	targetDriver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if err != nil {
		return "", err
	}
	if targetDriver == nil {
		fmt.Println("Driver not found")
		return "", nil
	}

	var driverConfig map[string]interface{}

	if len(args) == 4 {
		method := args[2]
		configValue := args[3]

		if method == "jsonfile" {
			fileContent, err := ioutil.ReadFile(configValue)
			if err != nil {
				return "", fmt.Errorf("Unable to read configuration file. %s", err.Error())
			}
			configValue = string(fileContent)
		}

		if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
			return "", fmt.Errorf("Invalid JSON format %s", err.Error())
		}
	} else if len(args) == 2 {

		configSchema, err := c.httpClient.GetDriverSchema(&operations.GetDriverSchemaParams{DriverID: *targetDriver.ID}, bearer)
		if err != nil {
			return "", err
		}

		configValue, err := c.schemaParser.ParseSchema(string(configSchema.Payload))
		if err != nil {
			return "", err
		}

		if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
			return "", fmt.Errorf("Invalid JSON format %s", err.Error())
		}
	}

	oldInstance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return "", err
	}
	if oldInstance == nil {
		fmt.Println("Driver instance not found")
		return "", nil
	}

	oldInstance.Configuration = driverConfig
	params := operations.NewUpdateDriverInstanceParams()
	params.DriverConfig = oldInstance
	params.DriverInstanceID = *oldInstance.ID
	params.DriverConfig.DriverID = *targetDriver.ID

	response, err := c.httpClient.UpdateDriverInstance(params, bearer)
	if err != nil {
		return "", err
	}

	return response.Payload.Name, nil
}

//List - lists existing instances for a specific driver
func (c *InstanceCommands) List(bearer swaggerclient.AuthInfoWriter, driverName string) ([]*models.DriverInstance, error) {
	targetDriver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if err != nil {
		return nil, err
	}
	if targetDriver == nil {
		fmt.Println("Driver not found")
		return nil, nil
	}

	params := operations.NewGetDriverInstancesParams()
	params.DriverID = *targetDriver.ID

	response, err := c.httpClient.GetDriverInstances(params, bearer)
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

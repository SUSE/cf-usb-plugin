package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/go-openapi/runtime"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
	"github.com/hpcloud/cf-plugin-usb/lib/schema"
)

//InstanceInterface exposes instances commands
type InstanceInterface interface {
	Create(runtime.ClientAuthInfoWriter, []string) (string, error)
	Delete(runtime.ClientAuthInfoWriter, string) (string, error)
	Update(runtime.ClientAuthInfoWriter, []string) (string, error)
	List(runtime.ClientAuthInfoWriter, string) ([]*models.DriverInstance, error)
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
func (c *InstanceCommands) Create(bearer runtime.ClientAuthInfoWriter, args []string) (string, error) {
	driverName := args[0]
	instanceName := args[1]
	targetUrl := args[2]

	targetDriver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if targetDriver == nil {
		return "", fmt.Errorf("Driver not found")
	}

	var driverConfig map[string]interface{}

	if len(args) == 5 {
		if args[3] == "-c" {
			configValue := args[4]

			if _, err := ioutil.ReadFile(configValue); err == nil {
				fileContent, err := ioutil.ReadFile(configValue)
				if err != nil {
					return "", fmt.Errorf("Unable to read configuration file. %s", err.Error())
				}
				configValue = string(fileContent)
			}

			if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
				return "", fmt.Errorf("Invalid JSON format %s", err.Error())
			}
		}
	} else if len(args) == 3 {
		configSchema, err := c.httpClient.GetDriverSchema(&operations.GetDriverSchemaParams{DriverID: targetDriver.ID}, bearer)
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
		Name:          &instanceName,
		DriverID:      &targetDriver.ID,
		Configuration: driverConfig,
		TargetURL:     targetUrl,
	}

	response, err := c.httpClient.CreateDriverInstance(&operations.CreateDriverInstanceParams{DriverInstance: &newDriver}, bearer)
	if err != nil {
		return "", err
	}

	return response.Payload.ID, nil
}

//Delete - deletes an existing driver instance
func (c *InstanceCommands) Delete(bearer runtime.ClientAuthInfoWriter, instanceName string) (string, error) {
	instance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return "", err
	}
	if instance == nil {
		return "", fmt.Errorf("Driver instance not found")
	}

	params := operations.NewDeleteDriverInstanceParams()
	params.DriverInstanceID = instance.ID

	_, err = c.httpClient.DeleteDriverInstance(params, bearer)
	if err != nil {
		return "", err
	}

	return instance.ID, nil
}

//Update - updates an existing driver instance
func (c *InstanceCommands) Update(bearer runtime.ClientAuthInfoWriter, args []string) (string, error) {
	instanceName := args[0]

	instance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return "", err
	}

	if instance.DriverID == nil {
		return "", fmt.Errorf("Empty driver id provided by cf-usb")
	}

	getDriverParams := operations.NewGetDriverParams()
	getDriverParams.DriverID = *instance.DriverID
	targetDriverResult, err := c.httpClient.GetDriver(getDriverParams, bearer)
	if err != nil {
		return "", err
	}

	targetDriver := targetDriverResult.Payload

	if targetDriver == nil {
		return "", fmt.Errorf("Driver not found")
	}

	var driverConfig map[string]interface{}

	if len(args) == 3 {
		if args[1] == "-c" {
			configValue := args[2]

			if _, err := ioutil.ReadFile(configValue); err == nil {
				fileContent, err := ioutil.ReadFile(configValue)
				if err != nil {
					return "", fmt.Errorf("Unable to read configuration file. %s", err.Error())
				}
				configValue = string(fileContent)
			}

			if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
				return "", fmt.Errorf("Invalid JSON format %s", err.Error())
			}
		}
	} else if len(args) == 1 {

		configSchema, err := c.httpClient.GetDriverSchema(&operations.GetDriverSchemaParams{DriverID: targetDriver.ID}, bearer)
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
		return "", fmt.Errorf("Driver instance not found")
	}

	oldInstance.Configuration = driverConfig
	params := operations.NewUpdateDriverInstanceParams()
	params.DriverConfig = oldInstance
	params.DriverInstanceID = oldInstance.ID
	params.DriverConfig.DriverID = &targetDriver.ID

	response, err := c.httpClient.UpdateDriverInstance(params, bearer)
	if err != nil {
		return "", err
	}

	return *response.Payload.Name, nil
}

type instanceSorter []*models.DriverInstance

func (a instanceSorter) Len() int           { return len(a) }
func (a instanceSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a instanceSorter) Less(i, j int) bool { return *a[i].Name < *a[j].Name }

//List - lists existing instances for a specific driver
func (c *InstanceCommands) List(bearer runtime.ClientAuthInfoWriter, driverName string) ([]*models.DriverInstance, error) {
	targetDriver, err := c.httpClient.GetDriverByName(bearer, driverName)
	if err != nil {
		return nil, err
	}
	if targetDriver == nil {
		return nil, fmt.Errorf("Driver not found")
	}

	params := operations.NewGetDriverInstancesParams()
	params.DriverID = targetDriver.ID

	response, err := c.httpClient.GetDriverInstances(params, bearer)
	if err != nil {
		return nil, err
	}
	sort.Sort(instanceSorter(response.Payload))
	return response.Payload, nil
}

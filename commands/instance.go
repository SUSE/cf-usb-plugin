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
	List(runtime.ClientAuthInfoWriter) ([]*models.DriverEndpoint, error)
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
	instanceName := args[0]
	targetUrl := args[1]
	authKey := args[2]

	var metadata models.EndpointMetadata

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

			if err := json.Unmarshal([]byte(configValue), &metadata); err != nil {
				return "", fmt.Errorf("Invalid JSON format %s", err.Error())
			}
		}
	}

	newDriver := models.DriverEndpoint{
		Name:              &instanceName,
		EndpointURL:       targetUrl,
		AuthenticationKey: authKey,
		Metadata:          &metadata,
	}

	response, err := c.httpClient.RegisterDriverEndpoint(&operations.RegisterDriverEndpointParams{DriverEndpoint: &newDriver}, bearer)
	if err != nil {
		return "", err
	}

	return response.Payload.ID, nil
}

//Delete - deletes an existing driver instance
func (c *InstanceCommands) Delete(bearer runtime.ClientAuthInfoWriter, instanceName string) (string, error) {
	instance, err := c.httpClient.GetDriverEndpointByName(instanceName, bearer)
	if err != nil {
		return "", err
	}
	if instance == nil {
		return "", fmt.Errorf("Driver instance not found")
	}

	params := operations.NewUnregisterDriverInstanceParams()
	params.DriverEndpointID = instance.ID

	_, err = c.httpClient.UnregisterDriverEndpoint(params, bearer)
	if err != nil {
		return "", err
	}

	return instance.ID, nil
}

//Update - updates an existing driver instance
func (c *InstanceCommands) Update(bearer runtime.ClientAuthInfoWriter, args []string) (string, error) {
	instanceName := args[0]

	var metadata models.EndpointMetadata

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

			if err := json.Unmarshal([]byte(configValue), &metadata); err != nil {
				return "", fmt.Errorf("Invalid JSON format %s", err.Error())
			}
		}
	}

	oldInstance, err := c.httpClient.GetDriverEndpointByName(instanceName, bearer)
	if err != nil {
		return "", err
	}
	if oldInstance == nil {
		return "", fmt.Errorf("Driver instance not found")
	}

	params := operations.NewUpdateDriverEndpointParams()
	params.DriverEndpointID = oldInstance.ID
	params.DriverEndpoint = &models.DriverEndpoint{}
	params.DriverEndpoint.Metadata = &metadata

	response, err := c.httpClient.UpdateDriverEndpoint(params, bearer)
	if err != nil {
		return "", err
	}

	return *response.Payload.Name, nil
}

type instanceSorter []*models.DriverEndpoint

func (a instanceSorter) Len() int           { return len(a) }
func (a instanceSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a instanceSorter) Less(i, j int) bool { return *a[i].Name < *a[j].Name }

//List - lists existing instances for a specific driver
func (c *InstanceCommands) List(bearer runtime.ClientAuthInfoWriter) ([]*models.DriverEndpoint, error) {

	params := operations.NewGetDriverEndpointsParams()
	response, err := c.httpClient.GetDriverEndpoints(params, bearer)
	if err != nil {
		return nil, err
	}
	sort.Sort(instanceSorter(response.Payload))
	return response.Payload, nil
}

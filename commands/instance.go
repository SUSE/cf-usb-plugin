package commands

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/go-openapi/runtime"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"

	httptransport "github.com/go-openapi/runtime/client"
)

//InstanceInterface exposes instances commands
type InstanceInterface interface {
	Create(string, string, string, *json.RawMessage) (string, error)
	Delete(string) (string, error)
	Update(string, *json.RawMessage) (string, error)
	List() ([]*models.DriverEndpoint, error)
}

//InstanceCommands struct
type InstanceCommands struct {
	httpClient lib.UsbClientInterface
	token      runtime.ClientAuthInfoWriter
}

//NewInstanceCommands - returns an InstanceCommands object
func NewInstanceCommands(httpClient lib.UsbClientInterface, bearer string) InstanceInterface {
	return &InstanceCommands{httpClient: httpClient, token: httptransport.BearerToken(bearer)}
}

//Create - creates a new driver instance
func (c *InstanceCommands) Create(instanceName, targetUrl, authKey string, metadata *json.RawMessage) (string, error) {

	newDriver := models.DriverEndpoint{
		Name:              &instanceName,
		EndpointURL:       targetUrl,
		AuthenticationKey: authKey,
	}

	if metadata != nil {
		var meta models.EndpointMetadata
		err := json.Unmarshal(*metadata, &meta)
		if err != nil {
			return "", err
		}

		newDriver.Metadata = &meta
	}

	response, err := c.httpClient.RegisterDriverEndpoint(&operations.RegisterDriverEndpointParams{DriverEndpoint: &newDriver}, c.token)
	if err != nil {
		return "", err
	}

	return response.Payload.ID, nil
}

//Delete - deletes an existing driver instance
func (c *InstanceCommands) Delete(instanceName string) (string, error) {
	instance, err := c.httpClient.GetDriverEndpointByName(instanceName, c.token)
	if err != nil {
		return "", err
	}
	if instance == nil {
		return "", fmt.Errorf("Driver instance not found")
	}

	params := operations.NewUnregisterDriverInstanceParams()
	params.DriverEndpointID = instance.ID

	_, err = c.httpClient.UnregisterDriverEndpoint(params, c.token)
	if err != nil {
		return "", err
	}

	return instance.ID, nil
}

//Update - updates an existing driver instance
func (c *InstanceCommands) Update(instanceName string, metadata *json.RawMessage) (string, error) {

	var meta models.EndpointMetadata
	err := json.Unmarshal(*metadata, &meta)
	if err != nil {
		return "", err
	}

	oldInstance, err := c.httpClient.GetDriverEndpointByName(instanceName, c.token)
	if err != nil {
		return "", err
	}
	if oldInstance == nil {
		return "", fmt.Errorf("Driver instance not found")
	}

	params := operations.NewUpdateDriverEndpointParams()
	params.DriverEndpointID = oldInstance.ID
	params.DriverEndpoint = &models.DriverEndpoint{}
	params.DriverEndpoint.Metadata = &meta

	response, err := c.httpClient.UpdateDriverEndpoint(params, c.token)
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
func (c *InstanceCommands) List() ([]*models.DriverEndpoint, error) {

	params := operations.NewGetDriverEndpointsParams()
	response, err := c.httpClient.GetDriverEndpoints(params, c.token)
	if err != nil {
		return nil, err
	}
	sort.Sort(instanceSorter(response.Payload))
	return response.Payload, nil
}

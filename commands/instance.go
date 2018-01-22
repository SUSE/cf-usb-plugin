package commands

import (
	"fmt"
	"sort"

	"github.com/SUSE/cf-usb-plugin/lib/client/operations"
	"github.com/go-openapi/runtime"

	"github.com/SUSE/cf-usb-plugin/lib"
	"github.com/SUSE/cf-usb-plugin/lib/models"

	httptransport "github.com/go-openapi/runtime/client"
)

//InstanceInterface exposes instances commands
type InstanceInterface interface {
	Create(instanceName, targetURL, authKey, CACert string, skipSSL *bool, metadata map[string]string) (string, error)
	Delete(string) (string, error)
	Update(string, string, string, map[string]string) (string, error)
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
func (c *InstanceCommands) Create(instanceName, targetURL, authKey string, caCert string, skipSSL *bool, metadata map[string]string) (string, error) {

	newDriver := models.DriverEndpoint{
		Name:              &instanceName,
		EndpointURL:       targetURL,
		AuthenticationKey: authKey,
		CaCertificate:     caCert,
		SkipSSLValidation: skipSSL,
	}

	if metadata != nil {
		newDriver.Metadata = metadata
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
func (c *InstanceCommands) Update(instanceName, targetUrl, authKey string, metadata map[string]string) (string, error) {

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
	if authKey != "" {
		params.DriverEndpoint.AuthenticationKey = authKey
	} else {
		params.DriverEndpoint.AuthenticationKey = oldInstance.AuthenticationKey
	}
	if targetUrl != "" {
		params.DriverEndpoint.EndpointURL = targetUrl
	} else {
		params.DriverEndpoint.EndpointURL = oldInstance.EndpointURL
	}

	params.DriverEndpoint.Metadata = metadata
	params.DriverEndpoint.Name = oldInstance.Name

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

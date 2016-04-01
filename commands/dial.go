package commands

import (
	"fmt"
	"sort"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

//DialInterface exposes list dials command
type DialInterface interface {
	List(swaggerclient.AuthInfoWriter, string) ([]*models.Dial, error)
}

//DialCommands struct
type DialCommands struct {
	httpClient lib.UsbClientInterface
}

//NewDialCommands returns a DialCommands object
func NewDialCommands(httpClient lib.UsbClientInterface) DialInterface {
	return &DialCommands{httpClient: httpClient}
}

type dialSorter []*models.Dial

func (a dialSorter) Len() int           { return len(a) }
func (a dialSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a dialSorter) Less(i, j int) bool { return *a[i].ID < *a[j].ID }

//List dials of an instance
func (c *DialCommands) List(bearer swaggerclient.AuthInfoWriter, instanceName string) ([]*models.Dial, error) {
	instance, err := c.httpClient.GetDriverInstanceByName(bearer, instanceName)
	if err != nil {
		return nil, err
	}
	if instance == nil {
		return nil, fmt.Errorf("Driver instance not found")
	}

	params := operations.NewGetAllDialsParams()
	params.DriverInstanceID = instance.ID

	response, err := c.httpClient.GetAllDials(params, bearer)
	if err != nil {
		return nil, err
	}

	sort.Sort(dialSorter(response.Payload))
	return response.Payload, nil
}

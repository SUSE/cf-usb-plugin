package commands

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	swaggerclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

func GetBearerToken(cliConnection plugin.CliConnection) (swaggerclient.AuthInfoWriter, error) {
	token, err := cliConnection.AccessToken()
	if err != nil {
		return nil, err
	}
	var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

	return bearer, nil
}

func getDriverByName(client *operations.Client, authHeader swaggerclient.AuthInfoWriter, driverName string) *models.Driver {
	ret, err := client.GetDrivers(&operations.GetDriversParams{}, authHeader)
	if err != nil {
		fmt.Println("ERROR - get driver by name:", err)
		return nil
	}

	var targetDriver *models.Driver

	for _, d := range ret.Payload {
		if d.Name == driverName {
			targetDriver = d
		}
	}

	return targetDriver
}

func getDriverInstanceByName(client *operations.Client, authHeader swaggerclient.AuthInfoWriter, driverInstanceName string) *models.DriverInstance {
	ret, err := client.GetDrivers(&operations.GetDriversParams{}, authHeader)
	if err != nil {
		fmt.Println("ERROR - get driver instance by name:", err)
		return nil
	}
	for _, d := range ret.Payload {
		for _, i := range d.DriverInstances {
			di, err := client.GetDriverInstance(&operations.GetDriverInstanceParams{DriverInstanceID: i}, authHeader)
			if err != nil {
				fmt.Println("ERROR - get driver instance by name:", err)
				return nil
			}
			if di.Payload.Name == driverInstanceName {
				return di.Payload
			}
		}
	}

	return nil
}

func getFileSha(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sha1 := sha1.New()
	_, err = io.Copy(sha1, f)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sha1.Sum(nil)), nil
}

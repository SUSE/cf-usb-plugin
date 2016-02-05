package main

import (
	"fmt"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	swaggerclient "github.com/go-swagger/go-swagger/client"
	"github.com/hpcloud/cf-plugin-usb/lib/models"
)

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

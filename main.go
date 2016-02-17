package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	swaggerclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/lib/client/operations"

	"github.com/hpcloud/cf-plugin-usb/lib/models"
	"github.com/hpcloud/cf-plugin-usb/lib/schema"
)

var target string

type UsbPlugin struct {
	ui         terminal.UI
	httpClient *operations.Client
}

func main() {
	plugin.Start(new(UsbPlugin))
}

func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	argLength := len(args)

	c.ui = terminal.NewUI(os.Stdin, terminal.NewTeePrinter())

	// except command to set target
	if !(args[1] == "target" && argLength == 3) {
		target, err := config.GetTarget()
		if target == "" {
			fmt.Println("Usb management target not set. Use cf usb target <usb-mgmt-endpoint> to set the target")
			return
		}
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		/*sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}*/
		u, err := url.Parse(target)
		if err != nil {
			fmt.Println("ERROR :", err)
			return
		}
		transport := httptransport.New(u.Host, "/", []string{u.Scheme})

		debug, _ := strconv.ParseBool(os.Getenv("CF_TRACE"))

		transport.Debug = debug

		c.httpClient = operations.New(transport, strfmt.Default)
	}

	switch args[1] {
	case "target":
		if argLength == 2 {
			target, err := config.GetTarget()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Println("Usb management target: " + target)
		} else if argLength == 3 {
			target = args[2]

			err := config.SetTarget(target)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Println("Usb management target set to: " + target)
		}
	case "info":
		token, err := cliConnection.AccessToken()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
		infoResp, err := c.httpClient.GetInfo(operations.NewGetInfoParams(), bearer)
		if err != nil {
			fmt.Println("ERROR:", err.Error())
			return
		}

		fmt.Println("info response: " + infoResp.Payload.Version)
	case "create-driver":
		if argLength == 4 || argLength == 5 {

			if argLength == 5 {
				if _, err := os.Stat(args[4]); err != nil {
					fmt.Println("ERROR:", err)
					return
				}
			}

			var driver models.Driver
			driver.DriverType = args[2]
			driver.Name = args[3]

			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			params := operations.NewCreateDriverParams()

			params.Driver = &driver

			response, err := c.httpClient.CreateDriver(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Driver created with ID:", *response.Payload.ID)

			filePath := ""
			if argLength == 5 {
				filePath = args[4]

				sha, err := getFileSha(filePath)
				if err != nil {
					fmt.Println("ERROR: ", err)
					return
				}

				file, err := os.Open(filePath)
				if err != nil {
					fmt.Println("ERROR: ", err)
					return
				}

				var uploadParams operations.UploadDriverParams

				uploadParams.DriverID = *response.Payload.ID
				uploadParams.File = *file
				uploadParams.Sha = sha

				_, err = c.httpClient.UploadDriver(&uploadParams, bearer)
				if err != nil {
					fmt.Println("ERROR:", err)
					return
				}
				fmt.Println("Uploaded driver bits from:", filePath)
			}
		} else {
			fmt.Println("ERROR: Invalid number of arguments")
			fmt.Println("Usage: create-driver [driver-type] [driver-name] [driver-bits-path]")
			return
		}
	case "create-instance":
		token, err := cliConnection.AccessToken()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
		if argLength == 6 {
			driverName := args[2]
			instanceName := args[3]
			method := args[4]
			configValue := args[5]

			if method == "configFile" {
				fileContent, err := ioutil.ReadFile(configValue)
				if err != nil {
					fmt.Println("ERROR - reading configuration file", err)
					return
				}
				configValue = string(fileContent)
			}

			fmt.Println(fmt.Sprintf("Creating instance %s for driver %s using %s with value %s", instanceName, driverName, method, configValue))

			var targetDriver *models.Driver = getDriverByName(c.httpClient, bearer, driverName)

			var driverConfig map[string]interface{}

			if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
				println("Invalid JSON format", err.Error())
			}

			newDriver := models.DriverInstance{
				Name:          instanceName,
				DriverID:      *targetDriver.ID,
				Configuration: driverConfig,
			}

			createRes, err := c.httpClient.CreateDriverInstance(&operations.CreateDriverInstanceParams{&newDriver}, bearer)
			if err != nil {
				fmt.Println("ERROR - create instance:", err)
				return
			}
			fmt.Println("New driver instance created. ID:" + *createRes.Payload.ID)
		} else {
			if argLength == 4 {
				driverName := args[2]
				instanceName := args[3]
				var targetDriver *models.Driver = getDriverByName(c.httpClient, bearer, driverName)

				if targetDriver == nil {
					println("Driver not found")
					return
				}
				fmt.Println(targetDriver)
				configSchema, err := c.httpClient.GetDriverSchema(&operations.GetDriverSchemaParams{DriverID: *targetDriver.ID}, bearer)
				if err != nil {
					fmt.Println("ERROR:", err)
					return
				}
				schemaParser := schema.NewSchemaParser(c.ui)

				configValue, err := schemaParser.ParseSchema(string(configSchema.Payload))
				if err != nil {
					fmt.Println("ERROR:", err)
					return
				}
				var driverConfig map[string]interface{}

				if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
					println("Invalid JSON format", err.Error())
				}

				newDriver := models.DriverInstance{
					Name:          instanceName,
					DriverID:      *targetDriver.ID,
					Configuration: driverConfig,
				}

				createRes, err := c.httpClient.CreateDriverInstance(&operations.CreateDriverInstanceParams{&newDriver}, bearer)
				if err != nil {
					fmt.Println("ERROR - create instance:", err)
					return
				}
				fmt.Println("New driver instance created. ID:" + *createRes.Payload.ID)
			} else {
				fmt.Println("Usage cf usb create-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
				return
			}
		}
	case "drivers":
		token, err := cliConnection.AccessToken()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
		resp, err := c.httpClient.GetDrivers(operations.NewGetDriversParams(), bearer)
		if err != nil {
			fmt.Println("ERROR:", err)
		}
		if resp != nil {
			table := terminal.NewTable(c.ui, []string{"Id", "Name", "Type"})
			for _, driver := range resp.Payload {
				table.Add(*driver.ID, driver.Name, driver.DriverType)
			}
			table.Print()
		}
	}
}

func (c *UsbPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cf-plugin-usb",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "usb",
				HelpText: "Usb plugin command's help text",
				UsageDetails: plugin.Usage{
					Usage: "cf usb",
				},
			},
			plugin.Command{
				Name:     "usb target",
				HelpText: "Gets or sets usb management endpoint",

				UsageDetails: plugin.Usage{
					Usage: "usb target <usb-mgmt-endpoint>\n   cf usb target <usb-mgmt-endpoint>",
				},
			},
			plugin.Command{
				Name:     "usb info",
				HelpText: "Usb plugin token command text",

				UsageDetails: plugin.Usage{
					Usage: "usb token\n   cf usb token",
				},
			},
			plugin.Command{
				Name:     "usb create-instance",
				HelpText: "Usb plugin create driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb create-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]",
				},
			},
			plugin.Command{
				Name:     "usb create-driver",
				HelpText: "Usb plugin create driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb create-instance [driverType] [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb drivers",
				HelpText: "List existing drivers",

				UsageDetails: plugin.Usage{
					Usage: "usb drivers\n   cf usb drivers",
				},
			},
		},
	}
}

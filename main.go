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

		fmt.Println("Broker API version: " + infoResp.Payload.BrokerAPIVersion)

		fmt.Println("USB version: " + infoResp.Payload.UsbVersion)
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
	case "delete-driver":
		if argLength == 3 {
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			driver := getDriverByName(c.httpClient, bearer, args[2])

			if driver == nil {
				fmt.Println("Driver not found")
				return
			}
			params := operations.NewDeleteDriverParams()
			params.DriverID = *driver.ID

			_, err = c.httpClient.DeleteDriver(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Driver deleted:", *driver.ID)
		} else {
			fmt.Println("Usage: delete-driver [driver-name]")
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
	case "delete-instance":
		if argLength == 3 {
			name := args[2]

			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			instance := getDriverInstanceByName(c.httpClient, bearer, name)

			params := operations.NewDeleteDriverInstanceParams()
			params.DriverInstanceID = *instance.ID

			response, err := c.httpClient.DeleteDriverInstance(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err.Error())
			}
			fmt.Println("Deleted driver instance response -", response)
		} else {
			fmt.Println("Usage cf usb delete-instance [instanceName]")
		}
	case "update-driver":
		if argLength == 4 {
			oldName := args[2]
			newName := args[3]

			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			driver := getDriverByName(c.httpClient, bearer, oldName)

			driver.Name = newName
			params := operations.NewUpdateDriverParams()
			params.DriverID = *driver.ID
			params.Driver = driver
			response, err := c.httpClient.UpdateDriver(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Driver updated:", response.Payload.Name)
		} else {
			fmt.Println("Usage: update-driver [old-driver-name] [new-driver-name]")
		}
	case "update-instance":
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
			var targetDriver *models.Driver = getDriverByName(c.httpClient, bearer, driverName)

			if targetDriver == nil {
				println("Driver not found")
				return
			}

			if method == "configFile" {
				fileContent, err := ioutil.ReadFile(configValue)
				if err != nil {
					fmt.Println("ERROR - reading configuration file", err)
					return
				}
				configValue = string(fileContent)
			}

			fmt.Println(fmt.Sprintf("Updating instance %s using %s with value %s", instanceName, method, configValue))

			oldinstance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
			if oldinstance == nil {
				println("Driver instance not found")
				return
			}
			var driverConfig map[string]interface{}

			if err := json.Unmarshal([]byte(configValue), &driverConfig); err != nil {
				println("Invalid JSON format", err.Error())
			}

			oldinstance.Configuration = driverConfig
			params := operations.NewUpdateDriverInstanceParams()
			params.DriverConfig = oldinstance
			params.DriverInstanceID = *oldinstance.ID
			params.DriverConfig.DriverID = *targetDriver.ID

			updateRes, err := c.httpClient.UpdateDriverInstance(params, bearer)
			if err != nil {
				fmt.Println("ERROR - update instance:", err)
				return
			}
			fmt.Println("Driver instance updated. ID:" + *updateRes.Payload.ID)
		} else {
			if argLength == 4 {
				driverName := args[2]
				instanceName := args[3]
				var targetDriver *models.Driver = getDriverByName(c.httpClient, bearer, driverName)

				if targetDriver == nil {
					println("Driver not found")
					return
				}

				oldinstance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
				if oldinstance == nil {
					println("Driver instance not found")
					return
				}

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

				oldinstance.Configuration = driverConfig
				params := operations.NewUpdateDriverInstanceParams()
				params.DriverConfig = oldinstance
				params.DriverInstanceID = *oldinstance.ID
				params.DriverConfig.DriverID = *targetDriver.ID

				updateRes, err := c.httpClient.UpdateDriverInstance(params, bearer)
				if err != nil {
					fmt.Println("ERROR - update instance:", err)
					return
				}
				fmt.Println("Driver instance updated. ID:" + *updateRes.Payload.ID)
			} else {
				fmt.Println("Usage: cf usb update-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
				return
			}
		}
	case "update-service":
		if argLength == 3 {
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			instanceName := args[2]

			instance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
			if instance == nil {
				fmt.Println("Driver instance not found")
				return
			}

			params := operations.NewUpdateServiceParams()
			params.ServiceID = *instance.Service

			var service models.Service
			service.DriverInstanceID = *instance.ID
			bindable := true
			bind := c.ui.Ask("Is service bindable?[Y/n]")
			if strings.ToLower(strings.Trim(bind, " ")) == "n" {
				bindable = false
			}
			service.Bindable = &bindable

			service.Name = c.ui.Ask("Service name")
			if service.Name == "" {
				fmt.Println("ERROR: Empty service name provided")
				return
			}
			desc := c.ui.Ask("Service description")
			if desc == "" {
				fmt.Println("ERROR: Empty service description provided")
				return
			}
			service.Description = &desc
			tags := c.ui.Ask("Tags (comma separated)")
			if tags == "" {
				fmt.Println("ERROR: Empty tags array provided")
				return
			}
			service.Tags = strings.Split(tags, ",")

			params.Service = &service

			response, err := c.httpClient.UpdateService(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Updated service with ID:", *response.Payload.ID)
		} else {
			fmt.Println("Usage: cf usb update-service [instanceName]")
			return
		}
	case "update-service-plan":
		if argLength == 4 {
			instanceName := args[2]
			planName := args[3]
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

			instance := getDriverInstanceByName(c.httpClient, bearer, instanceName)
			if instance == nil {
				fmt.Println("Driver instance not found")
				return
			}
			params := operations.NewGetAllDialsParams()
			params.DriverInstanceID = instance.ID

			dials, err := c.httpClient.GetAllDials(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			var oldPlan models.Plan

			for _, dial := range dials.Payload {
				plan, err := c.httpClient.GetServicePlan(&operations.GetServicePlanParams{PlanID: *dial.Plan}, bearer)
				if err != nil {
					fmt.Println("ERROR - get service plan", err)
				}
				if plan.Payload.Name == planName {
					oldPlan = *plan.Payload
					break
				}
			}
			if oldPlan.ID == nil {
				fmt.Println("Plan not found")
				return
			}
			planParams := operations.NewUpdateServicePlanParams()
			planParams.PlanID = *oldPlan.ID
			planParams.Plan = &oldPlan

			planParams.Plan.Name = c.ui.Ask("Plan name")
			if planParams.Plan.Name == "" {
				fmt.Println("Plan name cannot be empty")
				return
			}

			desc := c.ui.Ask("Plan description")
			if desc == "" {
				fmt.Println("Plan description cannot be empty")
				return
			}
			planParams.Plan.Description = &desc

			free := true
			f := c.ui.Ask("Is plan free?[Y/n]")
			if strings.ToLower(strings.Trim(f, " ")) == "n" {
				free = false
			}
			planParams.Plan.Free = &free

			response, err := c.httpClient.UpdateServicePlan(planParams, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			fmt.Println("Updated plan with ID:", response.Payload.ID)
		} else {
			fmt.Println("Usage: cf usb update-service-plan [instanceName] [planName]")
			return
		}
	case "dials":
		if argLength == 3 {
			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
			instance := getDriverInstanceByName(c.httpClient, bearer, args[2])
			if instance == nil {
				fmt.Println("Driver instance not found")
				return
			}
			params := operations.NewGetAllDialsParams()
			params.DriverInstanceID = instance.ID

			dials, err := c.httpClient.GetAllDials(params, bearer)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			for _, dial := range dials.Payload {
				fmt.Println("Dial configuration:\t", dial.Configuration)
				fmt.Println("Dial ID:\t\t", *dial.ID)
				fmt.Println("Plan ID:\t\t", *dial.Plan)

				plan, err := c.httpClient.GetServicePlan(&operations.GetServicePlanParams{PlanID: *dial.Plan}, bearer)
				if err != nil {
					fmt.Println("ERROR - getting plan", err)
				}
				fmt.Println("Plan:\t\t\t Name:", plan.Payload.Name, "; Description:", *plan.Payload.Description)
				fmt.Println("")
			}
		} else {
			fmt.Println("Usage: cf usb dials [instanceName]")
			return
		}
	case "instances":
		if argLength == 3 {

			token, err := cliConnection.AccessToken()
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			var bearer swaggerclient.AuthInfoWriter = httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))
			driver := getDriverByName(c.httpClient, bearer, args[2])
			if driver == nil {
				fmt.Println("ERROR: Driver not found")
			} else {
				params := operations.NewGetDriverInstancesParams()
				params.DriverID = *driver.ID
				response, err := c.httpClient.GetDriverInstances(params, bearer)
				if err != nil {
					fmt.Println("ERROR:", err)
					return
				}
				if response != nil {
					for _, di := range response.Payload {
						fmt.Println("Driver Instance Name:\t", di.Name)
						fmt.Println("Driver Instance Id:\t", *di.ID)
						fmt.Println("Configuration:\t\t", di.Configuration)
						fmt.Println("Dials:\t\t\t", len(di.Dials))
						service, err := c.httpClient.GetServiceByInstanceID(&operations.GetServiceByInstanceIDParams{DriverInstanceID: *di.ID}, bearer)
						if err != nil {
							fmt.Println("ERROR:", err)
						}

						fmt.Println("Service:\t\t", "Name:", service.Payload.Name, "; Bindable:", *service.Payload.Bindable, "; Tags:", service.Payload.Tags)
						fmt.Println("")
					}
				}
			}

		} else {
			fmt.Println("Usage cf usb instances [driverName]")
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
				Name:     "usb delete-instance",
				HelpText: "Usb plugin delete driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb delete-instance [instanceName]",
				},
			},
			plugin.Command{
				Name:     "usb create-driver",
				HelpText: "Usb plugin create driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb create-driver [driverType] [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb update-driver",
				HelpText: "Usb plugin update driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-driver [oldDriverName] [newDriverName]",
				},
			},
			plugin.Command{
				Name:     "usb update-instance",
				HelpText: "Usb plugin update driver instance command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-instance [driverName] [instanceName]  configValue/configFile [jsonValue/filePath]",
				},
			},
			plugin.Command{
				Name:     "usb update-service",
				HelpText: "Usb plugin update service command",

				UsageDetails: plugin.Usage{
					Usage: "usb update-service [instanceName]",
				},
			},
			plugin.Command{
				Name:     "usb delete-driver",
				HelpText: "Usb plugin delete driver command",

				UsageDetails: plugin.Usage{
					Usage: "usb delete-driver [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb drivers",
				HelpText: "List existing drivers",

				UsageDetails: plugin.Usage{
					Usage: "usb drivers\n   cf usb drivers",
				},
			},
			plugin.Command{
				Name:     "usb instances",
				HelpText: "List existing driver instances",

				UsageDetails: plugin.Usage{
					Usage: "usb instances  [driverName]",
				},
			},
			plugin.Command{
				Name:     "usb dials",
				HelpText: "List existing dials for instance",

				UsageDetails: plugin.Usage{
					Usage: "usb dials  [instanceName]",
				},
			},
		},
	}
}

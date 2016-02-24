package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	swaggerclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/hpcloud/cf-plugin-usb/commands"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/lib"
	"github.com/hpcloud/cf-plugin-usb/lib/schema"
)

var target string

//UsbPlugin struct
type UsbPlugin struct {
	argLength  int
	ui         terminal.UI
	token      swaggerclient.AuthInfoWriter
	httpClient lib.UsbClientInterface
}

func main() {
	plugin.Start(new(UsbPlugin))
}

//Run method called before each command
func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	c.argLength = len(args)

	config := config.NewConfig()

	c.ui = terminal.NewUI(os.Stdin, terminal.NewTeePrinter())

	bearer, err := commands.GetBearerToken(cliConnection)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	c.token = bearer

	// except command to set target
	if !(args[1] == "target" && c.argLength == 3) {
		var err error

		target, err = config.GetTarget()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if target == "" {
			fmt.Println("Usb management target not set. Use cf usb target <usb-mgmt-endpoint> to set the target")
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

		c.httpClient = lib.NewUsbClient(transport, strfmt.Default)
	}

	switch args[1] {
	case "target":
		c.TargetCommand(args, config)
	case "info":
		c.InfoCommand()
	case "create-driver":
		c.CreateDriverCommand(args)
	case "delete-driver":
		c.DeleteDriverCommand(args)
	case "create-instance":
		c.CreateInstanceCommand(args)
	case "delete-instance":
		c.DeleteInstanceCommand(args)
	case "update-driver":
		c.UpdateDriverCommand(args)
	case "update-instance":
		c.UpdateInstanceCommand(args)
	case "update-service":
		c.UpdateServiceCommand(args)
	case "dials":
		c.DialsCommand(args)
	case "instances":
		c.InstancesCommand(args)
	case "drivers":
		c.DriversCommand()
	}
}

//GetMetadata returns metadata for cf cli
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

//TargetCommand - gets or sets the target for the plugin
func (c *UsbPlugin) TargetCommand(args []string, config config.UsbConfigPluginInterface) {
	if c.argLength == 2 {
		var err error

		target, err = config.GetTarget()
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Usb management target: " + target)
	} else if c.argLength == 3 {
		target = args[2]
		err := config.SetTarget(target)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Usb management target set to: " + target)
	}
}

//InfoCommand - returns broker information
func (c *UsbPlugin) InfoCommand() {
	infoResp, err := commands.NewInfoCommands(c.httpClient).GetInfo(c.token)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Broker API version: " + infoResp.BrokerAPIVersion)

	fmt.Println("USB version: " + infoResp.UsbVersion)
}

//CreateDriverCommand - creates a new driver
func (c *UsbPlugin) CreateDriverCommand(args []string) {
	if c.argLength == 4 || c.argLength == 5 {
		createdDriverID, err := commands.NewDriverCommands(c.httpClient).Create(c.token, args[2:c.argLength])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Driver created with ID:", createdDriverID)
	} else {
		fmt.Println("ERROR: Invalid number of arguments")
		fmt.Println("Usage: cf usb create-driver [driver-type] [driver-name] [driver-bits-path]")
		return
	}
}

//DeleteDriverCommand - deletes an existing driver
func (c *UsbPlugin) DeleteDriverCommand(args []string) {
	if c.argLength == 3 {
		if c.ui.Confirm(fmt.Sprintf("Really delete the driver %v", args[2])) {
			deletedDriverID, err := commands.NewDriverCommands(c.httpClient).Delete(c.token, args[2])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if deletedDriverID == "" {
				fmt.Println("Driver not found")
			} else {
				fmt.Println("Driver deleted:", deletedDriverID)
			}
		}
	} else {
		fmt.Println("Usage: cf usb delete-driver [driver-name]")
	}
}

//CreateInstanceCommand - creates an instance of a driver
func (c *UsbPlugin) CreateInstanceCommand(args []string) {
	if c.argLength == 6 || c.argLength == 4 {
		schemaParser := schema.NewSchemaParser(c.ui)

		createdInstanceID, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Create(c.token, args[2:c.argLength])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if createdInstanceID != "" {
			fmt.Println("New driver instance created. ID:" + createdInstanceID)
		}

	} else {
		fmt.Println("Usage cf usb create-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
		return
	}
}

//DeleteInstanceCommand - deletes the instance of a driver
func (c *UsbPlugin) DeleteInstanceCommand(args []string) {
	if c.argLength == 3 {
		if c.ui.Confirm(fmt.Sprintf("Really delete the driver instance %v", args[2])) {
			schemaParser := schema.NewSchemaParser(c.ui)

			deletedInstanceID, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Delete(c.token, args[2])
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}
			if deletedInstanceID == "" {
				fmt.Println("Driver instance not found")
			} else {
				fmt.Println("Deleted driver instance:", deletedInstanceID)
			}
		}
	} else {
		fmt.Println("Usage cf usb delete-instance [instanceName]")
	}
}

//UpdateDriverCommand - allows user to change a drivers name
func (c *UsbPlugin) UpdateDriverCommand(args []string) {
	if c.argLength == 4 {
		updatedDriverName, err := commands.NewDriverCommands(c.httpClient).Update(c.token, args[2:c.argLength])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if updatedDriverName == "" {
			fmt.Println("Driver not found")
		} else {
			fmt.Println("Driver updated:", updatedDriverName)
		}
	} else {
		fmt.Println("Usage: cf usb update-driver [old-driver-name] [new-driver-name]")
	}
}

//UpdateInstanceCommand - allows user to update the instance of a driver
func (c *UsbPlugin) UpdateInstanceCommand(args []string) {
	if c.argLength == 6 || c.argLength == 4 {
		schemaParser := schema.NewSchemaParser(c.ui)

		updateInstanceName, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Update(c.token, args[2:c.argLength])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		if updateInstanceName != "" {
			fmt.Println("Driver instance updated:" + updateInstanceName)
		}
	} else {
		fmt.Println("Usage: cf usb update-instance [driverName] [instanceName] configValue/configFile [jsonValue/filePath]")
		return
	}
}

//UpdateServiceCommand - allows user to update the service provided by a driver instance
func (c *UsbPlugin) UpdateServiceCommand(args []string) {
	if c.argLength == 3 {
		instance, err := c.httpClient.GetDriverInstanceByName(c.token, args[2])
		if err != nil {
			fmt.Println("ERROR - get driver instance:", err)
			return
		}
		if instance == nil {
			fmt.Println("Driver instance not found")
			return
		}

		serviceCommand := commands.NewServiceCommands(c.httpClient)
		service, err := serviceCommand.GetServiceByDriverInstanceID(c.token, *instance.ID)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
		fmt.Println("service id:", *service.ID)
		service.DriverInstanceID = *instance.ID

		bindString := ""
		if *service.Bindable {
			bindString = "y"
		} else {
			bindString = "n"
		}

		bind := c.ui.Ask(fmt.Sprintf("Is service bindable ? (%s)", bindString))
		if bind != "" {
			bindable := true
			if strings.ToLower(strings.Trim(bind, " ")) == "n" {
				bindable = false
			}
			service.Bindable = &bindable
		}

		serviceName := c.ui.Ask(fmt.Sprintf("Service name (%s)", service.Name))
		if serviceName != "" {
			service.Name = serviceName
		}

		oldServiceDescription := ""
		if service.Description != nil {
			oldServiceDescription = fmt.Sprintf("(%s)", *service.Description)
		}

		serviceDesc := c.ui.Ask(fmt.Sprintf("Service description %s", oldServiceDescription))
		if serviceDesc != "" {
			service.Description = &serviceDesc
		}

		serviceTags := c.ui.Ask(fmt.Sprintf("Tags (comma separated) (%s)", strings.Join(service.Tags, ",")))
		if serviceTags != "" {
			service.Tags = strings.Split(serviceTags, ",")
		}

		serviceID, err := serviceCommand.Update(c.token, service)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		fmt.Println("Updated service with ID:", serviceID)
	} else {
		fmt.Println("Usage: cf usb update-service [instanceName]")
		return
	}
}

//DialsCommand - lists dials of a driver instance
func (c *UsbPlugin) DialsCommand(args []string) {
	if c.argLength == 3 {
		dials, err := commands.NewDialCommands(c.httpClient).List(c.token, args[2])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if dials != nil {
			planCommand := commands.NewPlanCommands(c.httpClient)

			for _, dial := range dials {
				fmt.Println("Dial configuration:\t", dial.Configuration)
				fmt.Println("Dial ID:\t\t", *dial.ID)
				fmt.Println("Plan ID:\t\t", *dial.Plan)

				plan, err := planCommand.GetPlanByID(c.token, *dial.Plan)
				if err != nil {
					fmt.Println("ERROR:", err)
				}
				fmt.Println("Plan:\t\t\t Name:", plan.Name, "; Description:", *plan.Description)
				fmt.Println("")
			}
		}
	} else {
		fmt.Println("Usage: cf usb dials [instanceName]")
		return
	}
}

//InstancesCommand - list instances of a driver
func (c *UsbPlugin) InstancesCommand(args []string) {
	if c.argLength == 3 {
		schemaParser := schema.NewSchemaParser(c.ui)
		instanceCommands := commands.NewInstanceCommands(c.httpClient, schemaParser)
		instances, err := instanceCommands.List(c.token, args[2])
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		if instances != nil {
			serviceCommand := commands.NewServiceCommands(c.httpClient)

			for _, di := range instances {
				fmt.Println("Driver Instance Name:\t", di.Name)
				fmt.Println("Driver Instance Id:\t", *di.ID)
				fmt.Println("Configuration:\t\t", di.Configuration)
				fmt.Println("Dials:\t\t\t", len(di.Dials))

				service, err := serviceCommand.GetServiceByDriverInstanceID(c.token, *di.ID)
				if err != nil {
					fmt.Println("ERROR:", err)
				}

				fmt.Println("Service:\t\t", "Name:", service.Name, "; Bindable:", *service.Bindable, "; Tags:", service.Tags)
				fmt.Println("")
			}
		}

	} else {
		fmt.Println("Usage cf usb instances [driverName]")
	}
}

//DriversCommand - list existing drivers
func (c *UsbPlugin) DriversCommand() {
	drivers, err := commands.NewDriverCommands(c.httpClient).List(c.token)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	if drivers != nil {
		table := terminal.NewTable(c.ui, []string{"Id", "Name", "Type"})
		for _, driver := range drivers {
			table.Add(*driver.ID, driver.Name, driver.DriverType)
		}
		table.Print()
	}
}

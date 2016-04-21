package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
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
	token      runtime.ClientAuthInfoWriter
	httpClient lib.UsbClientInterface
}

func main() {
	plugin.Start(new(UsbPlugin))
}

//Run method called before each command
func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	c.argLength = len(args)

	config := config.NewConfig()
	configFile := config.GetUsbConfigFile()

	if _, err := os.Stat(configFile); err != nil {
		_, err := cliConnection.HasAPIEndpoint()

		if err != nil {
			c.showFailed("The api endpoint doesn't exist")
			return
		}

		endpoint, err1 := cliConnection.ApiEndpoint()
		if err1 != nil {
			c.showFailed("Cannot connect to api endpoint")
			return
		}

		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			c.showFailed("Cannot create config file")
			return
		}

		usbendpoint := "usb." + strings.Replace(endpoint, "https://api.", "", 1)
		_, err2 := net.Dial("tcp", usbendpoint+":80")
		if err2 != nil {
			c.showFailed("Cannot connect to usb endpoint on port 80")
		}

		_, err3 := file.WriteString("{\"MgmtTarget\":\"http://" + usbendpoint + "\"}")

		if err3 != nil {
			c.showFailed("Error writing configuration to usb config file")
		}

		defer file.Close()

	}

	c.ui = terminal.NewUI(os.Stdin, terminal.NewTeePrinter())

	bearer, err := commands.GetBearerToken(cliConnection)
	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

	c.token = bearer

	if c.argLength == 1 || (c.argLength == 2 && args[1] == "help") {
		c.showCommandsWithHelpText()
		return
	}

	if c.argLength == 3 && args[1] == "help" {
		c.ui.Say(c.getUsage(args))
		return
	}

	// except command to set target
	if !(args[1] == "target" && c.argLength == 3) {
		var err error

		target, err = config.GetTarget()
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		/*sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}*/
		u, err := url.Parse(target)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
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
	case "rename-driver":
		c.RenameDriverCommand(args)
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
	default:
		fmt.Printf("'%s' is not a registered command. See 'cf usb help'", args[1])
		fmt.Println()
		return
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
				HelpText: "View command's help text",
				UsageDetails: plugin.Usage{
					Usage: "cf usb",
				},
			},
			plugin.Command{
				Name:     "usb target",
				HelpText: "Set or view target usb management endpoint api url",
				UsageDetails: plugin.Usage{
					Usage: "cf usb target [URL]",
				},
			},
			plugin.Command{
				Name:     "usb info",
				HelpText: "Show usb plugin info",
				UsageDetails: plugin.Usage{
					Usage: "cf usb info",
				},
			},
			plugin.Command{
				Name:     "usb create-instance",
				HelpText: "Create a driver instance",
				UsageDetails: plugin.Usage{
					Usage: `cf usb create-instance DRIVER_NAME INSTANCE_NAME [-c PARAMETERS_AS_JSON]

    Optionally provide driver-specific configuration parameters in a valid JSON object in-line:
    cf usb create-instance DRIVER_NAME INSTANCE_NAME -c '{"name":"value","name":"value"}'
	
    Optionally provide a file containing driver-specific configuration parameters in a valid JSON object.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb create-instance DRIVER_NAME INSTANCE_NAME -c PATH_TO_FILE	
					
EXAMPLE:
    cf usb create-instance mydriver myinstance (omit -c to configure driver instance interactively -- cf usb will prompt for config values)
    cf usb create-instance mydriver myinstance -c '{"host":"localhost","port":"1234","user":"username","password":"password"}'
    cf usb create-instance mydriver myinstance -c ~/workspace/tmp/mydriverinstance_config.json
	
OPTIONS:
    -c   Valid JSON object containing driver instance specific configuration parameters, provided in-line or in a file
`,
				},
			},
			plugin.Command{
				Name:     "usb delete-instance",
				HelpText: "Delete a driver instance",
				UsageDetails: plugin.Usage{
					Usage: "cf usb delete-instance INSTANCE_NAME",
				},
			},
			plugin.Command{
				Name:     "usb create-driver",
				HelpText: "Create a driver and upload driver bits",
				UsageDetails: plugin.Usage{
					Usage: "cf usb create-driver DRIVER_TYPE DRIVER_NAME DRIVER_BITS_PATH",
				},
			},
			plugin.Command{
				Name:     "usb rename-driver",
				HelpText: "Rename a driver",
				UsageDetails: plugin.Usage{
					Usage: "cf usb rename-driver OLD_DRIVER_NAME NEW_DRIVER_NAME",
				},
			},
			plugin.Command{
				Name:     "usb update-instance",
				HelpText: "Update a driver instance",
				UsageDetails: plugin.Usage{
					Usage: `cf usb update-instance INSTANCE_NAME [-c PARAMETERS_AS_JSON]

    Optionally provide driver-specific configuration parameters in a valid JSON object in-line:
    cf usb update-instance INSTANCE_NAME -c '{"name":"value","name":"value"}'
	
    Optionally provide a file containing driver-specific configuration parameters in a valid JSON object.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb update-instance INSTANCE_NAME -c PATH_TO_FILE	
					
EXAMPLE:
    cf usb update-instance myinstance (omit -c to configure driver instance interactively -- cf usb will prompt for config values)
    cf usb update-instance myinstance -c '{"host":"localhost","port":"1234","user":"username","password":"password"}'
    cf usb update-instance myinstance -c ~/workspace/tmp/myinstance_config.json
	
OPTIONS:
    -c   Valid JSON object containing driver instance specific configuration parameters, provided in-line or in a file
`,
				},
			},
			plugin.Command{
				Name:     "usb update-service",
				HelpText: "Update a service",
				UsageDetails: plugin.Usage{
					Usage: "cf usb update-service INSTANCE_NAME",
				},
			},
			plugin.Command{
				Name:     "usb delete-driver",
				HelpText: "Delete a driver",
				UsageDetails: plugin.Usage{
					Usage: "cf usb delete-driver DRIVER_NAME",
				},
			},
			plugin.Command{
				Name:     "usb drivers",
				HelpText: "List existing drivers",
				UsageDetails: plugin.Usage{
					Usage: "cf usb drivers",
				},
			},
			plugin.Command{
				Name:     "usb instances",
				HelpText: "List existing driver instances for a driver",
				UsageDetails: plugin.Usage{
					Usage: "cf usb instances DRIVER_NAME",
				},
			},
			plugin.Command{
				Name:     "usb dials",
				HelpText: "List existing dials for a driver instance",

				UsageDetails: plugin.Usage{
					Usage: "cf usb dials INSTANCE_NAME",
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
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		fmt.Println("Usb management target: " + target)
	} else if c.argLength == 3 {
		target = args[2]
		err := config.SetTarget(target)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		c.showOk(fmt.Sprint("Usb management target set to: " + target))
	}
}

//InfoCommand - returns broker information
func (c *UsbPlugin) InfoCommand() {
	infoResp, err := commands.NewInfoCommands(c.httpClient).GetInfo(c.token)
	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

	c.showOk("")
	fmt.Println("Broker API version: " + *infoResp.BrokerAPIVersion)
	fmt.Println("USB version: " + *infoResp.UsbVersion)
}

//CreateDriverCommand - creates a new driver
func (c *UsbPlugin) CreateDriverCommand(args []string) {
	if c.argLength == 4 {
		createdDriverID, err := commands.NewDriverCommands(c.httpClient).Create(c.token, args[2:c.argLength])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		c.showOk(fmt.Sprint("Driver created with ID:", createdDriverID))
	} else {
		c.showIncorrectUsage("Requires driver type, driver name, driver bits path as arguments\n", args)
	}
}

//DeleteDriverCommand - deletes an existing driver
func (c *UsbPlugin) DeleteDriverCommand(args []string) {
	if c.argLength == 3 {
		if c.ui.Confirm(fmt.Sprintf("Really delete the driver %v", args[2])) {
			deletedDriverID, err := commands.NewDriverCommands(c.httpClient).Delete(c.token, args[2])
			if err != nil {
				c.showFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if deletedDriverID == "" {
				c.showFailed("Driver not found")
			} else {
				c.showOk(fmt.Sprint("Driver deleted:", deletedDriverID))
			}
		}
	} else {
		c.showIncorrectUsage("Requires driver name as argument\n", args)
	}
}

//CreateInstanceCommand - creates an instance of a driver
func (c *UsbPlugin) CreateInstanceCommand(args []string) {
	if c.argLength == 7 || c.argLength == 5 {
		schemaParser := schema.NewSchemaParser(c.ui)
		createdInstanceID, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Create(c.token, args[2:c.argLength])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		if createdInstanceID != "" {
			c.showOk(fmt.Sprint("New driver instance created. ID:" + createdInstanceID))
		}

	} else {
		c.showIncorrectUsage("Requires driver name, instance name as arguments\n", args)
	}
}

//DeleteInstanceCommand - deletes the instance of a driver
func (c *UsbPlugin) DeleteInstanceCommand(args []string) {
	if c.argLength == 3 {
		if c.ui.Confirm(fmt.Sprintf("Really delete the driver instance %v", args[2])) {
			schemaParser := schema.NewSchemaParser(c.ui)

			deletedInstanceID, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Delete(c.token, args[2])
			if err != nil {
				c.showFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if deletedInstanceID == "" {
				c.showFailed("Driver instance not found")
			} else {
				c.showOk(fmt.Sprint("Deleted driver instance:", deletedInstanceID))
			}
		}
	} else {
		c.showIncorrectUsage("Requires instance name as argument\n", args)
	}
}

//RenameDriverCommand - allows user to change a drivers name
func (c *UsbPlugin) RenameDriverCommand(args []string) {
	if c.argLength == 4 {
		updatedDriverName, err := commands.NewDriverCommands(c.httpClient).Update(c.token, args[2:c.argLength])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		if updatedDriverName == "" {
			c.showFailed("Driver not found")
		} else {
			c.showOk(fmt.Sprint("Driver updated:", updatedDriverName))
		}
	} else {
		c.showIncorrectUsage("Requires old driver name, new driver name as arguments\n", args)
	}
}

//UpdateInstanceCommand - allows user to update the instance of a driver
func (c *UsbPlugin) UpdateInstanceCommand(args []string) {
	if c.argLength == 5 || c.argLength == 3 {
		schemaParser := schema.NewSchemaParser(c.ui)

		updateInstanceName, err := commands.NewInstanceCommands(c.httpClient, schemaParser).Update(c.token, args[2:c.argLength])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		if updateInstanceName != "" {
			c.showOk(fmt.Sprint("Driver instance updated:" + updateInstanceName))
		}
	} else {
		c.showIncorrectUsage("Requires instance name as argument\n", args)
	}
}

//UpdateServiceCommand - allows user to update the service provided by a driver instance
func (c *UsbPlugin) UpdateServiceCommand(args []string) {
	if c.argLength == 3 {
		instance, err := c.httpClient.GetDriverInstanceByName(c.token, args[2])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR - get driver instance:", err))
			return
		}
		if instance == nil {
			c.showFailed("Driver instance not found")
			return
		}

		service, err := c.httpClient.GetServiceByDriverInstanceID(c.token, instance.ID)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		fmt.Println("service id:", service.ID)
		service.DriverInstanceID = &instance.ID

		bindString := ""
		if service.Bindable {
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
			service.Bindable = bindable
		}

		serviceName := c.ui.Ask(fmt.Sprintf("Service name (%s)", service.Name))
		if serviceName != "" {
			service.Name = &serviceName
		}

		oldServiceDescription := ""
		if service.Description != "" {
			oldServiceDescription = fmt.Sprintf("(%s)", service.Description)
		}

		serviceDesc := c.ui.Ask(fmt.Sprintf("Service description %s", oldServiceDescription))
		if serviceDesc != "" {
			service.Description = serviceDesc
		}

		serviceTags := c.ui.Ask(fmt.Sprintf("Tags (comma separated) (%s)", strings.Join(service.Tags, ",")))
		if serviceTags != "" {
			service.Tags = strings.Split(serviceTags, ",")
		}

		serviceID, err := commands.NewServiceCommands(c.httpClient).Update(c.token, service)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		c.showOk(fmt.Sprint("Updated service with ID:", serviceID))
	} else {
		c.showIncorrectUsage("Requires instance name as argument\n", args)
	}
}

//DialsCommand - lists dials of a driver instance
func (c *UsbPlugin) DialsCommand(args []string) {
	if c.argLength == 3 {
		dials, err := commands.NewDialCommands(c.httpClient).List(c.token, args[2])
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		if dials != nil {
			c.showOk("")
			for _, dial := range dials {
				fmt.Println("Dial configuration:\t", dial.Configuration)
				fmt.Println("Dial ID:\t\t", dial.ID)
				fmt.Println("Plan ID:\t\t", dial.Plan)

				plan, err := c.httpClient.GetPlanByID(c.token, dial.Plan)
				if err != nil {
					c.showFailed(fmt.Sprint("ERROR:", err))
				}
				fmt.Println("Plan:\t\t\t Name:", plan.Name, "; Description:", plan.Description)
				fmt.Println("")
			}
		} else {
			c.showFailed("No dials found")
		}
	} else {
		c.showIncorrectUsage("Requires instance name as argument\n", args)
	}
}

//InstancesCommand - list instances of a driver
func (c *UsbPlugin) InstancesCommand(args []string) {
	schemaParser := schema.NewSchemaParser(c.ui)
	instanceCommands := commands.NewInstanceCommands(c.httpClient, schemaParser)
	instanceCount := 0

	drivers, err := commands.NewDriverCommands(c.httpClient).List(c.token)

	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}
	for _, driver := range drivers {
		instances, err := instanceCommands.List(c.token, *driver.Name)

		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}

		if instances != nil {
			for _, di := range instances {
				fmt.Println("Driver Instance Name:\t", *di.Name)
				fmt.Println("TARGET:\t\t", di.TargetURL)
				fmt.Println("Driver Instance Id:\t", di.ID)
				fmt.Println("Configuration:\t\t", di.Configuration)
				fmt.Println("Dials:\t\t\t", len(di.Dials))

				service, err := c.httpClient.GetServiceByDriverInstanceID(c.token, di.ID)

				if err != nil {
					c.showFailed(fmt.Sprint("ERROR:", err))
				}

				fmt.Println("Service:\t\t", "Name:", *service.Name, "; Bindable:", service.Bindable, "; Tags:", service.Tags)
				fmt.Println("")

				instanceCount++
			}
		}
	}

	if instanceCount == 0 {
		c.showFailed("No instances found")
	}
}

//DriversCommand - list existing drivers
func (c *UsbPlugin) DriversCommand() {
	drivers, err := commands.NewDriverCommands(c.httpClient).List(c.token)
	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

	driversCount := len(drivers)

	if driversCount > 0 {
		c.showOk("")
		table := terminal.NewTable(c.ui, []string{"Name", "Id", "Type"})
		for _, driver := range drivers {
			table.Add(*driver.Name, driver.ID, *driver.DriverType)
		}
		table.Print()
	} else {
		c.showFailed("No drivers found")
	}
}

func (c *UsbPlugin) showCommandsWithHelpText() {
	metadata := c.GetMetadata()
	for _, command := range metadata.Commands {
		fmt.Printf("%-25s %-50s\n", command.Name, command.HelpText)
	}
	return
}

func (c *UsbPlugin) getUsage(args []string) string {
	output := ""

	for _, cmd := range c.GetMetadata().Commands {
		command := args[1]
		if args[1] == "help" {
			command = args[2]
		}

		if cmd.Name == fmt.Sprintf("%s %s", args[0], command) {
			output = "NAME:\n    "
			output += fmt.Sprintf("%s - %s", cmd.Name, cmd.HelpText)
			output += "\n\nUSAGE:\n    "
			output += cmd.UsageDetails.Usage
			output += "\n"
		}
	}

	return output
}

func (c *UsbPlugin) showIncorrectUsage(message string, args []string) {
	c.ui.Failed(fmt.Sprintf("Incorrect Usage. %s\n", message) + c.getUsage(args))
}

func (c *UsbPlugin) showFailed(message string) {
	c.ui.Failed(fmt.Sprintf("%s\n", message))
}

func (c *UsbPlugin) showOk(message string) {
	c.ui.Say(terminal.SuccessColor("OK"))
	fmt.Printf("%s\n", message)
}

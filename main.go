package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/cf/trace"
	"github.com/cloudfoundry/cli/plugin"

	"github.com/hpcloud/cf-plugin-usb/commands"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/lib"
)

var target string

//UsbPlugin struct
type UsbPlugin struct {
	argLength  int
	ui         terminal.UI
	token      string
	httpClient lib.UsbClientInterface
}

func main() {
	plugin.Start(new(UsbPlugin))
}

//Run method called before each command
func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	c.argLength = len(args)

	traceEnv := os.Getenv("CF_TRACE")
	traceLogger := trace.NewLogger(Writer, false, traceEnv, "")

	config := config.NewConfig()
	configFile, err := config.GetUsbConfigFile()
	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

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

	c.ui = terminal.NewUI(os.Stdin, Writer, terminal.NewTeePrinter(Writer), traceLogger)

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

		debug, _ := strconv.ParseBool(os.Getenv("CF_TRACE"))
		c.httpClient = lib.NewUsbClient(u, debug)
	}

	switch args[1] {
	case "target":
		c.TargetCommand(args, config)
	case "info":
		c.InfoCommand()
	case "create-driver-endpoint":
		c.CreateInstanceCommand(args)
	case "delete-driver-endpoint":
		c.DeleteInstanceCommand(args)
	case "update-driver-endpoint":
		c.UpdateInstanceCommand(args)
	case "driver-endpoints":
		c.InstancesCommand(args)
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
				Name:     "usb create-driver-endpoint",
				HelpText: "Create a driver endpoint",
				UsageDetails: plugin.Usage{
					Usage: `cf usb create-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY [-c METADATA_AS_JSON]

    Optionally provide a file containing the driver endpoint metadata in a valid JSON object.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb create-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY -c PATH_TO_FILE	
					
EXAMPLE:
    cf usb create-driver-endpoint mydriver http://127.0.0.1:1234 authkey -c '{"display_name":"My Driver","image_url":"http://127.0.0.1:8080/image","long_description":"Long description","provider_display_name":"ProvidedName", "documentation_url":"http://127.0.0.1:8080/doc", "support_url":"http://127.0.0.1:8080/support"}'
    cf usb create-driver-endpoint mydriver http://127.0.0.1:1234 authkey -c ~/workspace/tmp/driver_metadata.json
	
OPTIONS:
    -c   Valid JSON object containing the driver endpoint metadata, provided in-line or in a file
`,
				},
			},
			plugin.Command{
				Name:     "usb delete-driver-endpoint",
				HelpText: "Delete a driver instance",
				UsageDetails: plugin.Usage{
					Usage: "cf usb delete-driver-endpoint NAME",
				},
			},
			plugin.Command{
				Name:     "usb update-driver-endpoint",
				HelpText: "Update a driver instance",
				UsageDetails: plugin.Usage{
					Usage: `cf usb update-driver-endpoint NAME [-c METADATA_AS_JSON]

    Optionally provide a file containing the driver endpoint metadata in a valid JSON object.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb update-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY -c PATH_TO_FILE	
					
EXAMPLE:
    cf usb update-driver-endpoint mydriver -c '{"display_name":"My Driver","image_url":"http://127.0.0.1:8080/image","long_description":"Long description","provider_display_name":"ProvidedName", "documentation_url":"http://127.0.0.1:8080/doc", "support_url":"http://127.0.0.1:8080/support"}'
    cf usb update-driver-endpoint mydriver -c ~/workspace/tmp/driver_metadata.json
	
OPTIONS:
    -c   Valid JSON object containing the driver endpoint metadata, provided in-line or in a file
`,
				},
			},
			plugin.Command{
				Name:     "usb driver-endpoints",
				HelpText: "List existing driver endpoints",
				UsageDetails: plugin.Usage{
					Usage: "cf usb driver-endpoints",
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
	infoResp, err := commands.NewInfoCommands(c.httpClient, c.token).GetInfo()
	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

	c.showOk("")
	fmt.Println("Broker API version: " + *infoResp.BrokerAPIVersion)
	fmt.Println("USB version: " + *infoResp.UsbVersion)
}

//CreateInstanceCommand - creates an instance of a driver
func (c *UsbPlugin) CreateInstanceCommand(args []string) {
	if c.argLength == 7 || c.argLength == 5 {
		instanceName := args[2]
		targetUrl := args[3]
		authKey := args[4]

		var rawMetadata *json.RawMessage

		if c.argLength == 7 {
			if args[5] == "-c" {
				configValue := args[6]

				if _, err := ioutil.ReadFile(configValue); err == nil {
					fileContent, err := ioutil.ReadFile(configValue)
					if err != nil {
						c.showFailed(fmt.Sprintf("Unable to read configuration file. %s", err.Error()))
					}
					configValue = string(fileContent)
				}
				if len(configValue) > 0 {
					meta := json.RawMessage(configValue)
					rawMetadata = &meta
				} else {
					rawMetadata = nil
				}
			}
		} else {
			rawMetadata = nil
		}

		createdInstanceID, err := commands.NewInstanceCommands(c.httpClient, c.token).Create(instanceName, targetUrl, authKey, rawMetadata)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		if createdInstanceID != "" {
			c.showOk(fmt.Sprint("New driver endpoint created. ID:" + createdInstanceID))
		}

	} else {
		c.showIncorrectUsage("Requires name, endpoint and auth key as arguments\n", args)
	}
}

//DeleteInstanceCommand - deletes the instance of a driver
func (c *UsbPlugin) DeleteInstanceCommand(args []string) {
	if c.argLength == 3 {
		if c.ui.Confirm(fmt.Sprintf("Really delete the driver endpoint %v", args[2])) {
			deletedInstanceID, err := commands.NewInstanceCommands(c.httpClient, c.token).Delete(args[2])
			if err != nil {
				c.showFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if deletedInstanceID == "" {
				c.showFailed("Driver endpoint not found")
			} else {
				c.showOk(fmt.Sprint("Deleted driver endpoint:", deletedInstanceID))
			}
		}
	} else {
		c.showIncorrectUsage("Requires endpoint name as argument\n", args)
	}
}

//UpdateInstanceCommand - allows user to update the instance of a driver
func (c *UsbPlugin) UpdateInstanceCommand(args []string) {
	if c.argLength == 5 {

		instanceName := args[2]

		var rawMetadata json.RawMessage

		if args[3] == "-c" {
			configValue := args[4]

			if _, err := ioutil.ReadFile(configValue); err == nil {
				fileContent, err := ioutil.ReadFile(configValue)
				if err != nil {
					c.showFailed(fmt.Sprintf("Unable to read configuration file. %s", err.Error()))
				}
				configValue = string(fileContent)
			}

			rawMetadata = json.RawMessage(configValue)
		}

		updateInstanceName, err := commands.NewInstanceCommands(c.httpClient, c.token).Update(instanceName, &rawMetadata)
		if err != nil {
			c.showFailed(fmt.Sprint("ERROR:", err))
			return
		}
		if updateInstanceName != "" {
			c.showOk(fmt.Sprint("Driver endpoint updated:" + updateInstanceName))
		}
	} else {
		c.showIncorrectUsage("Requires endpoint name as argument\n", args)
	}
}

//InstancesCommand - list endpoints
func (c *UsbPlugin) InstancesCommand(args []string) {
	instanceCommands := commands.NewInstanceCommands(c.httpClient, c.token)
	instanceCount := 0

	instances, err := instanceCommands.List()

	if err != nil {
		c.showFailed(fmt.Sprint("ERROR:", err))
		return
	}

	if instances != nil {
		for _, di := range instances {
			fmt.Println("Driver Endpoint Name:\t", *di.Name)
			fmt.Println("Endpoint URL:\t\t", di.EndpointURL)
			fmt.Println("Driver Endpoint Id:\t", di.ID)
			fmt.Println("Authentication Key:\t", di.AuthenticationKey)

			fmt.Println("Metadata:")
			fmt.Println("\tDisplayName:\t\t", di.Metadata.DisplayName)
			fmt.Println("\tDocumentationURL:\t", di.Metadata.DocumentationURL)
			fmt.Println("\tImageURL:\t\t", di.Metadata.ImageURL)
			fmt.Println("\tLongDescription:\t", di.Metadata.LongDescription)
			fmt.Println("\tProviderDisplayName:\t", di.Metadata.ProviderDisplayName)
			fmt.Println("\tSupportURL:\t\t", di.Metadata.SupportURL)

			fmt.Println()

			instanceCount++
		}
	}

	if instanceCount == 0 {
		c.showFailed("No instances found")
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

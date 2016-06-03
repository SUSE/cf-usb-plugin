package main

import (
	"fmt"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/cf/trace"
	"github.com/cloudfoundry/cli/plugin"

	"github.com/hpcloud/cf-plugin-usb/cmd"
	"github.com/hpcloud/cf-plugin-usb/commands"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/lib"
	usb "github.com/hpcloud/cf-plugin-usb/lib/plugin"
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
	traceLogger := trace.NewLogger(commands.Writer, false, traceEnv, "")

	config := config.NewConfig()
	configFile, err := config.GetUsbConfigFile()
	if err != nil {
		commands.ShowFailed(fmt.Sprint("ERROR:", err))
		return
	}

	if _, err := os.Stat(configFile); err != nil {
		_, err := cliConnection.HasAPIEndpoint()

		if err != nil {
			commands.ShowFailed("The api endpoint doesn't exist")
			return
		}

		endpoint, err1 := cliConnection.ApiEndpoint()
		if err1 != nil {
			commands.ShowFailed("Cannot connect to api endpoint")
			return
		}

		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			commands.ShowFailed("Cannot create config file")
			return
		}

		usbendpoint := "usb." + strings.Replace(endpoint, "https://api.", "", 1)
		_, err2 := net.Dial("tcp", usbendpoint+":80")
		if err2 != nil {
			commands.ShowFailed("Cannot connect to usb endpoint on port 80")
		}

		_, err3 := file.WriteString("{\"MgmtTarget\":\"http://" + usbendpoint + "\"}")

		if err3 != nil {
			commands.ShowFailed("Error writing configuration to usb config file")
		}

		defer file.Close()

	}

	c.ui = terminal.NewUI(os.Stdin, commands.Writer, terminal.NewTeePrinter(commands.Writer), traceLogger)

	bearer, err := commands.GetBearerToken(cliConnection)
	if err != nil {
		commands.ShowFailed(fmt.Sprint("ERROR:", err))
		return
	}

	c.token = bearer

	if c.argLength == 1 {
		c.showCommandsWithHelpText()
		return
	}

	// except command to set target
	if !(args[1] == "target" && c.argLength == 3) {
		var err error

		target, err = config.GetTarget()
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}

		/*sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}*/
		u, err := url.Parse(target)
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}

		debug, _ := strconv.ParseBool(os.Getenv("CF_TRACE"))
		c.httpClient = lib.NewUsbClient(u, debug)
	}

	usb.UsbClient.HttpClient = c.httpClient
	usb.UsbClient.Token = c.token
	usb.UsbClient.Commands = c.GetMetadata().Commands
	cmd.RootCmd.SetArgs(args[1:])
	cmd.Execute()
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
					Usage: `cf usb update-driver-endpoint NAME [-t ENDPOINT_URL] [-k AUTHENTICATION_KEY] [-c METADATA_AS_JSON]

    Optionally provide a file containing the driver endpoint metadata in a valid JSON object.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb update-driver-endpoint NAME -t ENDPOINT_URL -k AUTHENTICATION_KEY -c PATH_TO_FILE	
					
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

func (c *UsbPlugin) showCommandsWithHelpText() {
	metadata := c.GetMetadata()
	for _, command := range metadata.Commands {
		fmt.Printf("%-25s %-50s\n", command.Name, command.HelpText)
	}
	return
}

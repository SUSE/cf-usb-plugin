package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/cloudfoundry/cli/plugin"

	"github.com/SUSE/cf-usb-plugin/cmd"
	"github.com/SUSE/cf-usb-plugin/commands"
	"github.com/SUSE/cf-usb-plugin/config"
	"github.com/SUSE/cf-usb-plugin/lib"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"
)

var target string
var version string

//UsbPlugin struct
type UsbPlugin struct {
	argLength  int
	token      string
	httpClient lib.UsbClientInterface
}

func main() {
	plugin.Start(new(UsbPlugin))
}

//Run method called before each command
func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	c.argLength = len(args)

	if c.argLength < 2 {
		if c.argLength > 0 && strings.HasPrefix(args[0], "CLI-MESSAGE-") {
			// Internal CLI command (e.g. uninstall); don't show help text
		}
		return
	}

	config := config.NewConfig()
	configFile, err := config.GetUsbConfigFile()
	if err != nil {
		commands.ShowFailed(fmt.Sprint("ERROR:", err))
		return
	}

	if _, err := os.Stat(configFile); err != nil {
		_, err := cliConnection.HasAPIEndpoint()

		if err != nil {
			commands.ShowFailed(fmt.Sprintf("The api endpoint doesn't exist. Error: %s", err.Error()))
		}

		endpoint, err := cliConnection.ApiEndpoint()
		if err != nil {
			commands.ShowFailed(fmt.Sprintf("Cannot connect to api endpoint. Error: %s", err.Error()))
		}

		usbEndpoint := strings.Replace(endpoint, "api.", "usb.", 1)
		usbURL, err := url.Parse(usbEndpoint)
		if err != nil {
			commands.ShowFailed(fmt.Sprintf("The endpoint %s is not a valid URL. Error: %s", usbEndpoint, err.Error()))
		}

		_, err = net.Dial("tcp", fmt.Sprintf("%s:http", usbURL.Host))
		if err != nil {
			commands.ShowFailed(fmt.Sprintf("Cannot connect to usb endpoint %s on port 80. Error: %s", usbEndpoint, err.Error()))
		}

		file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
		defer file.Close()
		if err != nil {
			commands.ShowFailed(fmt.Sprintf("Cannot create config file. Error: %s", err.Error()))
		}

		_, err = file.WriteString(fmt.Sprintf(`{"MgmtTarget":"%s"}`, usbEndpoint))

		if err != nil {
			commands.ShowFailed(fmt.Sprintf("Error writing configuration to usb config file: %s", err.Error()))
		}

	}

	bearer, err := commands.GetBearerToken(cliConnection)
	if err != nil {
		commands.ShowFailed(fmt.Sprint("ERROR:", err))
	}

	c.token = bearer

	// except command to set target
	if !(args[0] == "usb-target" && c.argLength == 2) {
		var err error

		target, err = config.GetTarget()
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
		}

		/*sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}*/
		u, err := url.Parse(target)
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
		}

		debug, _ := strconv.ParseBool(os.Getenv("CF_TRACE"))
		sslDisabled, _ := cliConnection.IsSSLDisabled()
		c.httpClient = lib.NewUsbClient(u, sslDisabled, debug)
	}

	usb.UsbClient.HttpClient = c.httpClient
	usb.UsbClient.Token = c.token
	usb.UsbClient.Commands = c.GetMetadata().Commands
	cmd.RootCmd.SetArgs(args)
	cmd.Execute()
}

func getVersionInfo() [3]int {
	var tuple [3]int

	versionDigits := strings.SplitN(version, ".", 3)

	if len(versionDigits) != 3 {
		log.Fatal("Invalid version information")
	}

	var err error
	tuple[0], err = strconv.Atoi(versionDigits[0])
	if err != nil {
		log.Fatal("Invalid major version value")
	}

	tuple[1], err = strconv.Atoi(versionDigits[1])
	if err != nil {
		log.Fatal("Invalid minor version value")
	}

	buildDigits := strings.Split(versionDigits[2], "+")

	tuple[2], err = strconv.Atoi(buildDigits[0])

	if err != nil {
		log.Fatal("Invalid build version value")
	}

	return tuple
}

//GetMetadata returns metadata for cf cli
func (c *UsbPlugin) GetMetadata() plugin.PluginMetadata {
	versionDigits := getVersionInfo()

	return plugin.PluginMetadata{
		Name: "cf-usb-plugin",
		Version: plugin.VersionType{
			Major: versionDigits[0],
			Minor: versionDigits[1],
			Build: versionDigits[2],
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 14,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "usb-target",
				HelpText: "Set or view target usb management endpoint api url",
				UsageDetails: plugin.Usage{
					Usage: "cf usb-target [URL]",
				},
			},
			plugin.Command{
				Name:     "usb-info",
				HelpText: "Show usb plugin info",
				UsageDetails: plugin.Usage{
					Usage: "cf usb-info",
				},
			},
			plugin.Command{
				Name:     "usb-create-driver-endpoint",
				HelpText: "Create a driver endpoint",
				UsageDetails: plugin.Usage{
					Usage: `cf usb-create-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY [-c METADATA]

    Optionally provide a file containing the driver endpoint metadata in the following format mkey1:mval1;mkey2:mval2.
    The path to the parameters file can be an absolute or relative path to a file:
    cf usb create-driver-endpoint NAME ENDPOINT_URL AUTHENTICATION_KEY -c PATH_TO_FILE	
					
EXAMPLE:
    cf usb create-driver-endpoint mydriver http://127.0.0.1:1234 authkey -c 'mkey1:mval1;mkey2:mval2'
    cf usb create-driver-endpoint mydriver http://127.0.0.1:1234 authkey -c ~/workspace/tmp/driver_metadata.json
	
OPTIONS:
    -c   Valid JSON object containing the driver endpoint metadata, provided in-line or in a file
`,
				},
			},
			plugin.Command{
				Name:     "usb-delete-driver-endpoint",
				HelpText: "Delete a driver instance",
				UsageDetails: plugin.Usage{
					Usage: "cf usb-delete-driver-endpoint NAME",
				},
			},
			plugin.Command{
				Name:     "usb-update-driver-endpoint",
				HelpText: "Update a driver instance",
				UsageDetails: plugin.Usage{
					Usage: `cf usb-update-driver-endpoint NAME [-t ENDPOINT_URL] [-k AUTHENTICATION_KEY] [-c METADATA_AS_JSON]

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
				Name:     "usb-driver-endpoints",
				HelpText: "List existing driver endpoints",
				UsageDetails: plugin.Usage{
					Usage: "cf usb-driver-endpoints",
				},
			},
		},
	}
}

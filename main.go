package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/cf/terminal"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/hpcloud/cf-plugin-usb/config"
	"github.com/hpcloud/cf-plugin-usb/httpclient"
)

const brokerName string = "usb"

var target string

type UsbPlugin struct {
	ui         terminal.UI
	httpClient httpclient.HttpClient
}

func main() {
	plugin.Start(new(UsbPlugin))
}

func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if !usbBrokerExist(cliConnection) {
		fmt.Println("ERROR: No USB on this deployment")
		return
	}

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
		}

		sslDisabled, err := cliConnection.IsSSLDisabled()
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		c.httpClient = httpclient.NewHttpClient(target, sslDisabled)
	}

	switch args[1] {
	case "target":
		fmt.Println("Running the usb target command")

		if argLength == 2 {
			target, err := config.GetTarget()
			if err != nil {
				fmt.Println("ERROR:", err)
			}

			fmt.Printf("Usb management target: %s", target)
		} else if argLength == 3 {
			target = args[2]

			err := config.SetTarget(target)
			if err != nil {
				fmt.Println("ERROR:", err)
				return
			}

			fmt.Printf("Usb management target set to: %s", target)
		}
	case "info":
		fmt.Println("Running the usb info command")

		token, err := cliConnection.AccessToken()
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		getInfoReq := httpclient.Request{Verb: "GET", ApiUrl: "/info", Authorization: token, StatusCode: 200}

		getInfoResp, err := c.httpClient.Request(getInfoReq)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Printf("result: %s", string(getInfoResp))
	case "drivers":
		fmt.Println("Not implemented")

		token, err := cliConnection.AccessToken()
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Printf("token: %s", token)

		// ask user to add an input
		//value := c.ui.Ask("Value")
	}

	fmt.Println(terminal.ColorizeBold("OK", 32))
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
			Major: 6,
			Minor: 7,
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
				Name:     "usb drivers",
				HelpText: "List existing drivers",

				UsageDetails: plugin.Usage{
					Usage: "usb drivers\n   cf usb drivers",
				},
			},
		},
	}
}

func usbBrokerExist(cliConnection plugin.CliConnection) bool {
	brokers, err := cliConnection.CliCommandWithoutTerminalOutput("service-brokers")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	for _, a := range brokers {
		fields := strings.Fields(a)
		if fields[0] == brokerName {
			return true
		}
	}

	return false
}

package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/hpcloud/cf-plugin-usb/httpclient"
)

const brokerName string = "usb"

type UsbPlugin struct {
	httpClient httpclient.HttpClient
}

func main() {
	plugin.Start(new(UsbPlugin))
}

func (c *UsbPlugin) Run(cliConnection plugin.CliConnection, args []string) {

	brokerMgmtUrl, err := serviceBrokerUrl(cliConnection)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	sslDisabled, err := cliConnection.IsSSLDisabled()
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	c.httpClient = httpclient.NewHttpClient(brokerMgmtUrl, sslDisabled)

	// Ensure that we called the command usb info command
	switch args[1] {
	case "info":
		fmt.Println("Running the usb plugin info command")

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
				Name:     "usb info",
				HelpText: "Usb plugin token command text",

				// UsageDetails is optional
				// It is used to show help of usage of each command
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

func serviceBrokerUrl(cliConnection plugin.CliConnection) (string, error) {
	brokers, err := cliConnection.CliCommandWithoutTerminalOutput("service-brokers")
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	for _, a := range brokers {
		fields := strings.Fields(a)
		if fields[0] == brokerName {
			return fields[1], nil
		}
	}
	return "", errors.New("No such broker")
}

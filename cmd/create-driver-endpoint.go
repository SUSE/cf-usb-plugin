package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hpcloud/cf-plugin-usb/commands"
	usb "github.com/hpcloud/cf-plugin-usb/lib/plugin"

	"github.com/spf13/cobra"
)

// create-driver-endpointCmd represents the create-driver-endpoint command
var createDriverEndpointCmd = &cobra.Command{
	Use:   "create-driver-endpoint",
	Short: "Creates a new driver endpoint",
	Long:  `Creates a new driver endpoint registration in the usb`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			instanceName := args[0]
			targetUrl := args[1]
			authKey := args[2]

			var rawMetadata *json.RawMessage

			if len(configJson) > 0 {
				configValue := configJson

				if _, err := ioutil.ReadFile(configValue); err == nil {
					fileContent, err := ioutil.ReadFile(configValue)
					if err != nil {
						commands.ShowFailed(fmt.Sprintf("Unable to read configuration file. %s", err.Error()))
					}
					configValue = string(fileContent)
				}
				if len(configValue) > 0 {
					meta := json.RawMessage(configValue)
					rawMetadata = &meta
				} else {
					rawMetadata = nil
				}
			} else {
				rawMetadata = nil
			}

			createdInstanceID, err := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).Create(instanceName, targetUrl, authKey, caCert, &skipSSL, rawMetadata)
			if err != nil {
				commands.ShowFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if createdInstanceID != "" {
				commands.ShowOK(fmt.Sprint("New driver endpoint created. ID:" + createdInstanceID))
			}

		} else {
			commands.ShowIncorrectUsage("Requires name, endpoint and auth key as arguments\n", []string{"usb create-driver-endpoint"})
		}
	},
}

func init() {
	createDriverEndpointCmd.Flags().StringVarP(&configJson, "configuration", "c", "", "metadata configuration")
	createDriverEndpointCmd.Flags().StringVarP(&caCert, "ca-certificate", "x", "", "CA Certificate for TLS")
	createDriverEndpointCmd.Flags().BoolVarP(&skipSSL, "skip-csm-ssl-validation", "k", false, "Skip SSL Validation")
	RootCmd.AddCommand(createDriverEndpointCmd)
}

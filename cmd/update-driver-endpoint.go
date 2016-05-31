package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/hpcloud/cf-plugin-usb/commands"
	usb "github.com/hpcloud/cf-plugin-usb/lib/plugin"

	"github.com/spf13/cobra"
)

// update-driver-endpointCmd represents the update-driver-endpoint command
var updateDriverEndpointCmd = &cobra.Command{
	Use:   "update-driver-endpoint",
	Short: "Update driver information",
	Long:  `Updates the registered driver endpoint definitions in the usb`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 1 {
			instanceName := args[0]
			targetUrl := target
			authKey := key

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
			updateInstanceName, err := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).Update(instanceName, targetUrl, authKey, rawMetadata)
			if err != nil {
				commands.ShowFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if updateInstanceName != "" {
				commands.ShowOK(fmt.Sprint("Driver endpoint updated:" + updateInstanceName))
			}
		} else {
			commands.ShowIncorrectUsage("Requires endpoint name as argument\n", []string{"usb update-driver-endpoint"})
		}
	},
}

func init() {
	updateDriverEndpointCmd.Flags().StringVarP(&configJson, "configuration", "c", "", "metadata configuration")
	updateDriverEndpointCmd.Flags().StringVarP(&target, "target", "t", "", "driver endpoint target url")
	updateDriverEndpointCmd.Flags().StringVarP(&key, "authkey", "k", "", "authorization key")

	RootCmd.AddCommand(updateDriverEndpointCmd)
}

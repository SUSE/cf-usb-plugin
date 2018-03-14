package cmd

import (
	"fmt"
	"strings"

	"github.com/SUSE/cf-usb-plugin/commands"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"

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
			targetURL := target
			authKey := key

			metadata := make(map[string]string)
			rows := strings.Split(configJSON, ";")
			for _, row := range rows {
				key := strings.Split(row, ":")[0]
				value := strings.Split(row, ":")[1]
				metadata[key] = value
			}

			updateInstanceName, err := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).Update(instanceName, targetURL, authKey, metadata)
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
	updateDriverEndpointCmd.Flags().StringVarP(&configJSON, "configuration", "c", "", "metadata configuration")
	updateDriverEndpointCmd.Flags().StringVarP(&target, "target", "t", "", "driver endpoint target url")
	updateDriverEndpointCmd.Flags().StringVarP(&key, "authkey", "k", "", "authorization key")

	RootCmd.AddCommand(updateDriverEndpointCmd)
}

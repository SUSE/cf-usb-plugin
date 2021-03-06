package cmd

import (
	"fmt"
	"strings"

	"github.com/SUSE/cf-usb-plugin/commands"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"

	"github.com/spf13/cobra"
)

// create-driver-endpointCmd represents the create-driver-endpoint command
var createDriverEndpointCmd = &cobra.Command{
	Use:   "usb-create-driver-endpoint",
	Short: "Creates a new driver endpoint",
	Long:  `Creates a new driver endpoint registration in the usb`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			instanceName := args[0]
			targetURL := args[1]
			authKey := args[2]

			metadata := make(map[string]string)

			rows := strings.Split(configJSON, ";")
			for _, row := range rows {
				key := strings.Split(row, ":")[0]
				value := strings.Split(row, ":")[1]
				metadata[key] = value
			}

			createdInstanceID, err := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).Create(instanceName, targetURL, authKey, caCert, &skipSSL, metadata)
			if err != nil {
				commands.ShowFailed(fmt.Sprint("ERROR:", err))
				return
			}
			if createdInstanceID != "" {
				commands.ShowOK(fmt.Sprint("New driver endpoint created. ID:" + createdInstanceID))
			}

		} else {
			commands.ShowIncorrectUsage("Requires name, endpoint and auth key as arguments\n", []string{"usb-create-driver-endpoint"})
		}
	},
}

func init() {
	createDriverEndpointCmd.Flags().StringVarP(&configJSON, "configuration", "c", "", "metadata configuration")
	createDriverEndpointCmd.Flags().StringVarP(&caCert, "ca-certificate", "x", "", "CA Certificate for TLS")
	createDriverEndpointCmd.Flags().BoolVarP(&skipSSL, "skip-csm-ssl-validation", "k", false, "Skip SSL Validation")
	RootCmd.AddCommand(createDriverEndpointCmd)
}

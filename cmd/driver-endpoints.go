package cmd

import (
	"fmt"

	"github.com/SUSE/cf-usb-plugin/commands"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"

	"github.com/spf13/cobra"
)

// driver-endpointsCmd represents the driver-endpoints command
var driverEndpointsCmd = &cobra.Command{
	Use:   "driver-endpoints",
	Short: "Lists driver endpoints",
	Long:  `Shows a list of all the available driver endpoints that were registered with the usb`,
	Run: func(cmd *cobra.Command, args []string) {

		instanceCommands := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token)
		instanceCount := 0

		instances, err := instanceCommands.List()

		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}

		if instances != nil {
			for _, di := range instances {
				fmt.Println("Driver Endpoint Name:\t", *di.Name)
				fmt.Println("Endpoint URL:\t\t", di.EndpointURL)
				fmt.Println("Driver Endpoint Id:\t", di.ID)
				fmt.Println("Authentication Key:\t", di.AuthenticationKey)

				fmt.Println("Metadata:")
				for _, item := range di.Metadata {
					fmt.Println(item)
				}

				fmt.Println()

				instanceCount++
			}
		}

		if instanceCount == 0 {
			commands.ShowFailed("No instances found")
		}
	},
}

func init() {
	RootCmd.AddCommand(driverEndpointsCmd)
}

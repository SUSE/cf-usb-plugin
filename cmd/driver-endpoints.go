package cmd

import (
	"fmt"

	"github.com/hpcloud/cf-plugin-usb/commands"
	usb "github.com/hpcloud/cf-plugin-usb/lib/plugin"

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
			commands.ShowFailed("No instances found")
		}
	},
}

func init() {
	RootCmd.AddCommand(driverEndpointsCmd)
}

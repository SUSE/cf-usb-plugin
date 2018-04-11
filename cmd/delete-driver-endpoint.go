package cmd

import (
	"fmt"

	"github.com/SUSE/cf-usb-plugin/commands"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"

	"github.com/spf13/cobra"
)

// delete-driver-endpointCmd represents the delete-driver-endpoint command
var deleteDriverEndpointCmd = &cobra.Command{
	Use:   "usb-delete-driver-endpoint",
	Short: "Deletes a driver endpoint",
	Long:  `Removes a driver endpoint from the usb`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			if commands.Confirm(fmt.Sprintf("Really delete the driver endpoint %v", args[0])) {
				deletedInstanceID, err := commands.NewInstanceCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).Delete(args[0])
				if err != nil {
					commands.ShowFailed(fmt.Sprint("ERROR:", err))
					return
				}
				if deletedInstanceID == "" {
					commands.ShowFailed("Driver endpoint not found")
				} else {
					commands.ShowOK(fmt.Sprint("Deleted driver endpoint:", deletedInstanceID))
				}
			}
		} else {
			commands.ShowIncorrectUsage("Requires endpoint name as argument\n", []string{"usb-delete-driver-endpoint"})
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteDriverEndpointCmd)
}

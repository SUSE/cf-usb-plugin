package cmd

import (
	"fmt"

	"github.com/SUSE/cf-usb-plugin/commands"
	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get broker api information",
	Long:  `Get broker api version and api version information from the targeted usb broker`,
	Run: func(cmd *cobra.Command, args []string) {
		infoResp, err := commands.NewInfoCommands(usb.UsbClient.HttpClient, usb.UsbClient.Token).GetInfo()
		if err != nil {
			commands.ShowFailed(fmt.Sprint("ERROR:", err))
			return
		}

		commands.ShowOK("")
		fmt.Println("Broker API version: " + *infoResp.BrokerAPIVersion)
		fmt.Println("USB version: " + *infoResp.UsbVersion)
	},
}

func init() {
	RootCmd.AddCommand(infoCmd)
}

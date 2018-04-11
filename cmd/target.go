package cmd

import (
	"fmt"
	"github.com/SUSE/cf-usb-plugin/commands"
	"github.com/SUSE/cf-usb-plugin/config"
	"github.com/spf13/cobra"
)

// targetCmd represents the target command
var targetCmd = &cobra.Command{
	Use:   "usb-target",
	Short: "Gets or sets the target endpoint for the usb plugin",
	Long:  `Shows the current target or sets a new usb management url as a target for the plugin`,
	Run: func(cmd *cobra.Command, args []string) {
		configuration := config.NewConfig()
		if len(args) == 0 {
			var err error
			target, err := configuration.GetTarget()
			if err != nil {
				commands.ShowFailed(fmt.Sprint("ERROR:", err))
				return
			}

			fmt.Println("Usb management target: " + target)
		} else if len(args) == 1 {
			target := args[0]
			err := configuration.SetTarget(target)
			if err != nil {
				commands.ShowFailed(fmt.Sprint("ERROR:", err))
				return
			}

			commands.ShowOK(fmt.Sprint("Usb management target set to: " + target))
		}
	},
}

func init() {
	RootCmd.AddCommand(targetCmd)
}

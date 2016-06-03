// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/hpcloud/cf-plugin-usb/commands"
	usb "github.com/hpcloud/cf-plugin-usb/lib/plugin"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Shows help information",
	Long:  `Show command help information`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(commands.GetUsage([]string{"usb " + args[0]}))
		} else {
			for _, command := range usb.UsbClient.Commands {
				fmt.Printf("%-25s %-50s\n", command.Name, command.HelpText)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(helpCmd)
}

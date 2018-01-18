package commands

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"

	usb "github.com/SUSE/cf-usb-plugin/lib/plugin"

	"github.com/cloudfoundry/cli/plugin"
)

//GetBearerToken - returns token from cf cli
func GetBearerToken(cliConnection plugin.CliConnection) (string, error) {
	token, err := cliConnection.AccessToken()
	if err != nil {
		return "", err
	}
	return strings.Replace(token, "bearer ", "", -1), nil
}

func getFileSha(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	sha1 := sha1.New()
	_, err = io.Copy(sha1, f)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(sha1.Sum(nil)), nil
}

// GetUsage returns the help string for the application
func GetUsage(args []string) string {
	output := ""
	for _, cmd := range usb.UsbClient.Commands {
		command := args[0]
		if args[0] == "help" {
			command = args[1]
		}

		if cmd.Name == command {
			output = "NAME:\n    "
			output += fmt.Sprintf("%s - %s", cmd.Name, cmd.HelpText)
			output += "\n\nUSAGE:\n    "
			output += cmd.UsageDetails.Usage
			output += "\n"
		}
	}

	return output
}

// ShowIncorrectUsage lets the user know there is an operator error
func ShowIncorrectUsage(message string, args []string) {
	ShowFailed(fmt.Sprintf("Incorrect Usage. %s\n", message) + GetUsage(args))
}

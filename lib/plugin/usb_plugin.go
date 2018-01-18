package plugin

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/SUSE/cf-usb-plugin/lib"
)

var UsbClient struct {
	Token      string
	HttpClient lib.UsbClientInterface
	Commands   []plugin.Command
}

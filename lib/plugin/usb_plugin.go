package plugin

import (
	"github.com/cloudfoundry/cli/plugin"
	"github.com/hpcloud/cf-plugin-usb/lib"
)

var UsbClient struct {
	Token      string
	HttpClient lib.UsbClientInterface
	Commands   []plugin.Command
}

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/config_helpers"
)

//UsbConfigPluginInterface exposes config commands
type UsbConfigPluginInterface interface {
	GetTarget() (string, error)
	SetTarget(string) error
}

//UsbConfigPlugin struct
type UsbConfigPlugin struct {
}

//NewConfig returns a UsbConfigPlugin object
func NewConfig() UsbConfigPluginInterface {
	return &UsbConfigPlugin{}
}

//GetTarget returns selected usb target
func (*UsbConfigPlugin) GetTarget() (target string, err error) {
	jsonConf, err := ioutil.ReadFile(getUsbConfigFile())
	if err != nil {
		return "", err
	}

	config := NewData()

	err = config.ReadConfigData(jsonConf)
	if err != nil {
		return "", err
	}

	return config.MgmtTarget, nil
}

//SetTarget saves the target information in config file
func (*UsbConfigPlugin) SetTarget(target string) (err error) {
	if !strings.Contains(target, "http") {
		target = fmt.Sprintf("http://%[1]s", target)
	}

	file, err := os.OpenFile(getUsbConfigFile(), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer file.Close()

	config := NewData()
	config.MgmtTarget = target

	output, err := config.SetConfigData()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getUsbConfigFile(), output, 0600)
	if err != nil {
		return err
	}

	return nil
}

func getUsbConfigFile() string {
	return filepath.Join(filepath.Dir(config_helpers.DefaultFilePath()), "usb-config.json")
}

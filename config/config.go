package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/confighelpers"
)

//UsbConfigPluginInterface exposes config commands
type UsbConfigPluginInterface interface {
	GetTarget() (string, error)
	SetTarget(string) error
	GetUsbConfigFile() (string, error)
}

//UsbConfigPlugin struct
type UsbConfigPlugin struct {
}

//NewConfig returns a UsbConfigPlugin object
func NewConfig() UsbConfigPluginInterface {
	return &UsbConfigPlugin{}
}

//GetTarget returns selected usb target
func (configfile *UsbConfigPlugin) GetTarget() (target string, err error) {
	configFilePath, err := configfile.GetUsbConfigFile()
	if err != nil {
		return "", err
	}
	if _, err = os.Stat(configFilePath); err != nil {
		return "", errors.New("Usb management target not set. Use cf usb target <usb-mgmt-endpoint> to set the target")
	}

	jsonConf, err := ioutil.ReadFile(configFilePath)
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
func (configfile *UsbConfigPlugin) SetTarget(target string) (err error) {
	if !strings.Contains(target, "http") {
		target = fmt.Sprintf("http://%[1]s", target)
	}

	configFilePath, err := configfile.GetUsbConfigFile()
	if err != nil {
		return err
	}

	file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0755)
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

	err = ioutil.WriteFile(configFilePath, output, 0600)
	if err != nil {
		return err
	}

	return nil
}

//GetUsbConfigFile returns the path to the usb config file
func (*UsbConfigPlugin) GetUsbConfigFile() (string, error) {
	defaultPath, err := confighelpers.DefaultFilePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(filepath.Dir(defaultPath), "usb-config.json"), nil
}

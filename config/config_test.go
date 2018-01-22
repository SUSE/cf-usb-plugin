package config_test

import (
	"github.com/SUSE/cf-usb-plugin/config/fakes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SetTarget(t *testing.T) {
	assert := assert.New(t)

	config := new(fakes.FakeUsbConfigPluginInterface)

	err := config.SetTarget("http://testTarget")
	assert.NoError(err)
}

func Test_GetTarget(t *testing.T) {
	assert := assert.New(t)

	config := new(fakes.FakeUsbConfigPluginInterface)

	config.GetTargetReturns("http://testTarget", nil)

	target, err := config.GetTarget()

	assert.NoError(err)
	assert.Equal(target, "http://testTarget")
}

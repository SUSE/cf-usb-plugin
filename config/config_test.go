package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"os"
	"path/filepath"
)

func Test_SetTarget(t *testing.T) {
	assert := assert.New(t)

	homedir := os.Getenv("HOME")

	err := os.MkdirAll(filepath.Join(homedir,".cf"),0755)

	assert.NoError(err)

	config := NewConfig()

	err = config.SetTarget("http://testTarget")
	assert.NoError(err)
}

func Test_GetTarget(t *testing.T) {
	assert := assert.New(t)

	config := NewConfig()

	target, err := config.GetTarget()

	assert.NoError(err)
	assert.Equal(target, "http://testTarget")
}

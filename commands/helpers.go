package commands

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"os"
	"strings"

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

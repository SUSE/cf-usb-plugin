package commands

import (
	"crypto/sha1"
	"encoding/base64"
	httptransport "github.com/go-openapi/runtime/client"
	"io"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/go-openapi/runtime"
)

//GetBearerToken - returns token from cf cli
func GetBearerToken(cliConnection plugin.CliConnection) (runtime.ClientAuthInfoWriter, error) {
	token, err := cliConnection.AccessToken()
	if err != nil {
		return nil, err
	}
	bearer := httptransport.BearerToken(strings.Replace(token, "bearer ", "", -1))

	return bearer, nil
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

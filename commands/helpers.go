package commands

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"os"
	"strings"

	"github.com/cloudfoundry/cli/plugin"
	swaggerclient "github.com/go-swagger/go-swagger/client"
	httptransport "github.com/go-swagger/go-swagger/httpkit/client"
)

//GetBearerToken - returns token from cf cli
func GetBearerToken(cliConnection plugin.CliConnection) (swaggerclient.AuthInfoWriter, error) {
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

package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"strings"

	"github.com/cloudfoundry/cli/cf/configuration/core_config"
	"github.com/cloudfoundry/cli/cf/errors"
	. "github.com/cloudfoundry/cli/cf/i18n"
	"github.com/cloudfoundry/cli/cf/net"
)

type CurlRepository interface {
	Request(method, path, header, body string) (resHeaders, resBody string, apiErr error)
}

type CloudControllerCurlRepository struct {
	config  core_config.Reader
	gateway net.Gateway
}

func NewCloudControllerCurlRepository(config core_config.Reader, gateway net.Gateway) (repo CloudControllerCurlRepository) {
	repo.config = config
	repo.gateway = gateway
	return
}

func (repo CloudControllerCurlRepository) Request(method, path, headerString, body string) (resHeaders, resBody string, err error) {
	url := fmt.Sprintf("%s/%s", repo.config.ApiEndpoint(), strings.TrimLeft(path, "/"))

	req, err := repo.gateway.NewRequest(method, url, repo.config.AccessToken(), strings.NewReader(body))
	if err != nil {
		return
	}

	err = mergeHeaders(req.HttpReq.Header, headerString)
	if err != nil {
		err = fmt.Errorf("%s: %s", T("Error parsing headers"), err.Error())
		return
	}

	res, err := repo.gateway.PerformRequest(req)

	if _, ok := err.(errors.HttpError); ok {
		err = nil
	}

	if err != nil {
		return
	}
	defer res.Body.Close()

	headerBytes, _ := httputil.DumpResponse(res, false)
	resHeaders = string(headerBytes)

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("%s: %s", T("Error reading response"), err.Error())
	}
	resBody = string(bytes)

	return
}

func mergeHeaders(destination http.Header, headerString string) (err error) {
	headerString = strings.TrimSpace(headerString)
	headerString += "\n\n"
	headerReader := bufio.NewReader(strings.NewReader(headerString))
	headers, err := textproto.NewReader(headerReader).ReadMIMEHeader()
	if err != nil {
		return
	}

	for key, values := range headers {
		destination.Del(key)
		for _, value := range values {
			destination.Add(key, value)
		}
	}

	return
}

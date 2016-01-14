package httpclient

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient interface {
	Request(request Request) ([]byte, error)
}

type Request struct {
	Verb          string
	ApiUrl        string
	Body          io.ReadSeeker
	Authorization string
	StatusCode    int
}

type httpClient struct {
	skipTslValidation bool
	endpoint          string
}

func NewHttpClient(endpoint string, skipTslValidation bool) HttpClient {
	return &httpClient{
		skipTslValidation: skipTslValidation,
		endpoint:          endpoint,
	}
}

func (client *httpClient) Request(request Request) ([]byte, error) {
	httpResponse, err := client.httpRequest(request)
	if err != nil {
		return nil, err
	}

	return httpResponse, nil
}

func (client *httpClient) httpRequest(req Request) ([]byte, error) {
	request, err := http.NewRequest(req.Verb, client.endpoint+req.ApiUrl, req.Body)
	if err != nil {
		return nil, errors.New("Error building request")
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Authorization", req.Authorization)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: client.skipTslValidation},
	}
	httpClient := &http.Client{Transport: tr}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != req.StatusCode {
		return nil, errors.New(fmt.Sprintf("ERROR: status code: %d, body: %s", response.StatusCode, responseBody))
	}

	return responseBody, nil
}

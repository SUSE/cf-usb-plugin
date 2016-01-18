package info

import (
	"github.com/hpcloud/cf-plugin-usb/httpclient"
)

type InfoInterface interface {
	GetInfo() (string, error)
}

type Info struct {
	client httpclient.HttpClient
	token  string
}

func NewInfo(client httpclient.HttpClient, token string) InfoInterface {
	return &Info{
		client: client,
		token:  token,
	}
}

func (c *Info) GetInfo() (string, error) {
	getInfoReq := httpclient.Request{Verb: "GET", ApiUrl: "/info", Authorization: c.token, StatusCode: 200}

	getInfoResp, err := c.client.Request(getInfoReq)
	if err != nil {
		return "", err
	}

	return string(getInfoResp), nil
}

package fflogs

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	BaseUrlDefault = "https://www.fflogs.com:443/v1"
)

type Client struct {
	apiKey     string
	httpClient *http.Client // default: http.DefaultClient
	baseUrl    url.URL      // default: fflogs.BaseUrlDefault

	Raw *Raw
}

type ClientOpt struct {
	ApiKey     string
	HttpClient *http.Client // default: http.DefaultClient
	BaseUrl    string       // default: fflogs.BaseUrlDefault
}

func NewClient(opt *ClientOpt) (*Client, error) {
	if opt.ApiKey == "" {
		return nil, errors.New("ApiKey cannot be empty")
	}

	if opt.HttpClient == nil {
		opt.HttpClient = http.DefaultClient
	}
	if opt.BaseUrl == "" {
		opt.BaseUrl = BaseUrlDefault
	}

	u, err := url.Parse(opt.BaseUrl)
	if err != nil {
		return nil, err
	}

	client := &Client{
		httpClient: opt.HttpClient,
		apiKey:     opt.ApiKey,
		baseUrl:    *u,
	}
	client.Raw = &Raw{
		client: client,
	}

	return client, nil
}

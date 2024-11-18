package client

import (
	"crypto/tls"

	"github.com/go-resty/resty/v2"
)

func New(baseUrl string) *resty.Client {
	client := resty.New().SetBaseURL(baseUrl).SetTLSClientConfig(
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)

	return client
}

func NewWithRetry(baseUrl string) *resty.Client {
	client := New(baseUrl)

	client.SetRetryCount(3)
	client.SetRetryWaitTime(5000)
	client.SetRetryMaxWaitTime(20000)

	return client
}

package client

import (
	"crypto/tls"
	"fmt"
	"packages/server/config"

	"resty.dev/v3"
)

func GetAmosRequest(auth string) *resty.Request {
	c := resty.New()

	c.SetBaseURL(fmt.Sprintf("%s:%d/service/", config.AMOS.Host, config.AMOS.Port))
	c.SetHeader("Authorization", fmt.Sprintf("Basic %s", auth))
	c.SetHeader("Content-Type", "application/xml")

	c.SetTLSClientConfig(
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)

	return c.R()
}

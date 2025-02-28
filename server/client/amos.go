package client

import (
	"crypto/tls"
	"fmt"
	"packages/server/config"

	"resty.dev/v3"
)

func GetAmosRequest() *resty.Request {
	c := resty.New()

	c.SetBaseURL(fmt.Sprintf("%s:%d/service/", config.AMOS.Host, config.AMOS.Port))
	c.SetBasicAuth(config.AMOS.Auth.ID, config.AMOS.Auth.Password)
	c.SetHeader("Content-Type", "application/xml")

	c.SetTLSClientConfig(
		&tls.Config{
			InsecureSkipVerify: true,
		},
	)

	return c.R()
}

package amos

import (
	"errors"
	"packages/src/config"
)

type WebServiceAuth struct {
	Id       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CheckWebServiceAuth(webService *WebServiceAuth) error {
	if webService.Id != config.Amos.WebserviceId {
		return errors.New("invalid amos web-service id")
	}

	if webService.Password != config.Amos.WebservicePassword {
		return errors.New("invalid amos web-service password")
	}

	return nil
}

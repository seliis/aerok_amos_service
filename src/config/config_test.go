package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	path := os.Getenv("USERPROFILE") + "/mis/dev/aerok_amos_service/config.toml"

	if err := Load(path); err != nil {
		t.Error(err)
	}

	t.Log("Server.Port:", Server.Port)
	t.Log("Server.Cert.Key:", Server.Cert.Key)
	t.Log("Server.Cert.Pem:", Server.Cert.Pem)
	t.Log("Server.Service.DataSource:", Server.Service.DataSource)
	t.Log("Server.Service.BaseCurrency:", Server.Service.BaseCurrency)
	t.Log("Amos.BaseUrl:", Amos.BaseUrl)
	t.Log("Amos.WebserviceId:", Amos.WebserviceId)
	t.Log("Amos.WebservicePassword:", Amos.WebservicePassword)
	t.Log("Amos.ImportCurrency.Endpoint:", Amos.ImportCurrency.Endpoint)
	t.Log("Amos.ImportCurrency.BaseCurrency:", Amos.ImportCurrency.BaseCurrency)
	t.Log("Amos.ImportCurrency.Applicables:", Amos.ImportCurrency.Applicables)
}

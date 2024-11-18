package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

var Server server
var Amos amos

type config struct {
	Server server
	Amos   amos
}

type server struct {
	Port int
	Cert struct {
		Key string
		Pem string
	}
	Service struct {
		BaseCurrency     string `toml:"base_currency"`
		KoreaEximAuthKey string `toml:"korea_exim_auth_key"`
	}
}

type amos struct {
	BaseUrl            string `toml:"base_url"`
	WebserviceId       string `toml:"webservice_id"`
	WebservicePassword string `toml:"webservice_password"`
	ImportCurrency     struct {
		Endpoint     string
		BaseCurrency string `toml:"base_currency"`
		Applicables  []struct {
			CurrencyCode string `toml:"currency_code"`
			CurrencyName string `toml:"currency_name"`
		}
	} `toml:"import_currency"`
}

func Load(path string) error {
	var config config

	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := toml.Unmarshal(bytes, &config); err != nil {
		return err
	}

	Server = config.Server
	Amos = config.Amos

	return nil
}

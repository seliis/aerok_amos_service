package config

import (
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"
)

var (
	once     sync.Once
	instance *_Config
	Server   *_Server
	Amos     *_Amos
)

type _Config struct {
	Server _Server `toml:"server"`
	Amos   _Amos   `toml:"amos"`
}

type _Server struct {
	Static string `toml:"static"`
	Debug  bool   `toml:"debug"`
	Port   int    `toml:"port"`
	Cert   struct {
		Key string `toml:"key"`
		Pem string `toml:"pem"`
	} `toml:"cert"`
	Service struct {
		BaseCurrency     string `toml:"base_currency"`
		KoreaEximAuthKey string `toml:"korea_exim_auth_key"`
	}
}

type _Amos struct {
	BaseUrl    string `toml:"base_url"`
	WebService struct {
		Id       string `toml:"id"`
		Password string `toml:"password"`
	} `toml:"web_service"`
	ImportCurrency struct {
		Endpoint     string `toml:"endpoint"`
		BaseCurrency string `toml:"base_currency"`
		Currencies   []struct {
			Code string `toml:"code" json:"code"`
			Name string `toml:"name" json:"name"`
		}
	} `toml:"import_currency"`
	FutureFlights struct {
		Endpoint string `toml:"endpoint"`
	} `toml:"future_flights"`
}

func Load(path string) error {
	var err error

	once.Do(func() {
		instance, err = read(path)
		if err != nil {
			return
		}

		Server = &instance.Server
		Amos = &instance.Amos
	})

	return err
}

func read(path string) (*_Config, error) {
	var config _Config

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

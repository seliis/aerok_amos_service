package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	Server    *_Server
	Database  *_Database
	KoreaExim *_KoreaExim
	AMOS      *_AMOS
)

type _Config struct {
	Server    _Server    `toml:"server"`
	Database  _Database  `toml:"database"`
	KoreaExim _KoreaExim `toml:"korea-exim"`
	AMOS      _AMOS      `toml:"amos"`
}

type _Server struct {
	Protocol int `toml:"protocol"`
	Port     int `toml:"port"`
}

type _Database struct {
	URL string `toml:"url"`
}

type _KoreaExim struct {
	Host string `toml:"host"`
	Path string `toml:"path"`
	Auth string `toml:"auth"`
}

type _AMOS struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Auth struct {
		ID       string `toml:"id"`
		Password string `toml:"password"`
	} `toml:"auth"`
	Services struct {
		ImportCurrency struct {
			EndPoint         string `toml:"end_point"`
			AmosCurrencyCode string `toml:"amos_currency_code"`
		} `toml:"import_currency"`
		TransferFutureFlights struct {
			EndPoint string `toml:"end_point"`
		} `toml:"transfer_future_flights"`
	} `toml:"services"`
}

func Load(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var config _Config

	if err := toml.Unmarshal(bytes, &config); err != nil {
		return err
	}

	Server = &config.Server
	Database = &config.Database
	KoreaExim = &config.KoreaExim
	AMOS = &config.AMOS

	return nil
}

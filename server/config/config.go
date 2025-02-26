package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

var (
	Server _Server
)

type _Server struct {
	Port     int       `toml:"port"`
	Database _Database `toml:"database"`
}

type _Database struct {
	URL string `toml:"url"`
}

func Load(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := toml.Unmarshal(bytes, &Server); err != nil {
		return err
	}

	return nil
}

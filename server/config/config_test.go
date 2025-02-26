package config_test

import (
	"packages/server/config"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := config.Load("../../settings.toml"); err != nil {
		t.Fatal(err)
	}

	t.Logf("Server.Port: %d", config.Server.Port)
	t.Logf("Server.Database.URL: %s", config.Server.Database.URL)
}

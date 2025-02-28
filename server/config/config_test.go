package config_test

import (
	"packages/server/config"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := config.Load("../../settings.toml"); err != nil {
		t.Error(err)
	}

	t.Logf("Server.Port: %d", config.Server.Port)
	t.Logf("Database.URL: %s", config.Database.URL)
	t.Logf("AMOS.ImportCurrency.Path: %s", config.AMOS.Services.ImportCurrency.EndPoint)
}

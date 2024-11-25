package config_test

import (
	"encoding/json"
	"os"
	"packages/src/config"
	"testing"
)

func TestLoad(t *testing.T) {
	path := os.Getenv("USERPROFILE") + "/mis/dev/aerok_amos_service/config.toml"

	if err := config.Load(path); err != nil {
		t.Error(err)
	}

	bytes, err := json.MarshalIndent(config.Amos, "", "  ")
	if err != nil {
		t.Error(err)
	}

	t.Log(string(bytes))
}

package amos_test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"packages/server/amos"
	"packages/server/config"
	"packages/server/database"
	"packages/server/internal/services"
	"testing"
)

func TestImportCurrency(t *testing.T) {
	if err := config.Load("../../settings.toml"); err != nil {
		t.Error(err)
	}

	if err := database.Connect("file:../../prisma/database.db"); err != nil {
		t.Error(err)
	}

	s := services.NewExchangeRateService()

	t.Run("ImportCurrency", func(t *testing.T) {
		exchangeRates, err := s.GetExchangeRates(context.Background(), "2025-02-27")
		if err != nil {
			t.Error(err)
		}

		data, err := amos.NewImportCurrency(exchangeRates)
		if err != nil {
			t.Error(err)
		}

		json, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			t.Error(err)
		}

		t.Log(string(json))
	})

	t.Run("Push", func(t *testing.T) {
		exchangeRates, err := s.GetExchangeRates(context.Background(), "2025-02-27")
		if err != nil {
			t.Error(err)
		}

		data, err := amos.NewImportCurrency(exchangeRates)
		if err != nil {
			t.Error(err)
		}

		token := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.AMOS.Auth.ID, config.AMOS.Auth.Password))))

		if err := data.Push(token); err != nil {
			t.Error(err)
		}
	})

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Error(err)
		}
	}()
}

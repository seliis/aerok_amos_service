package amos_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"packages/server/config"
	"packages/server/database"
	"packages/server/internal/services"
	"testing"
)

func TestTransferFutureFlights(t *testing.T) {
	if err := config.Load("../../settings.toml"); err != nil {
		t.Error(err)
	}

	if err := database.Connect("file:../../prisma/database.db"); err != nil {
		t.Error(err)
	}

	t.Run("TransferFutureFlights", func(t *testing.T) {
		s := services.NewAmosService()

		token := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.AMOS.Auth.ID, config.AMOS.Auth.Password))))

		if err := s.TransferFutureFlights(context.Background(), token); err != nil {
			t.Error(err)
		}
	})

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Error(err)
		}
	}()
}

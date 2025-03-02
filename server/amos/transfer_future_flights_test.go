package amos_test

import (
	"context"
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

		if err := s.TransferFutureFlights(context.Background(), "2025-03-01"); err != nil {
			t.Error(err)
		}
	})

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Error(err)
		}
	}()
}

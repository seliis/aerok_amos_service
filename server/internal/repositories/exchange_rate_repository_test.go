package repositories_test

import (
	"context"
	"packages/server/database"
	"packages/server/internal/repositories"
	"testing"
)

func TestExchangeRateRepository(t *testing.T) {
	if err := database.Connect("file:../../../prisma/database.db"); err != nil {
		t.Error(err)
	}

	r := repositories.NewExchangeRateRepository()

	t.Run("ListByDate", func(t *testing.T) {
		arr, err := r.ListByDate(context.Background(), "2025-02-27")
		if err != nil {
			t.Error(err)
		}

		t.Log("ListByDate: ", len(arr))
	})

	t.Run("GetExchangeRates", func(t *testing.T) {
		arr, err := r.GetExchangeRates(context.Background(), "2025-02-26")
		if err != nil {
			t.Error(err)
		}

		for _, v := range arr {
			t.Logf("%+v", v)
		}
	})

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Error(err)
		}
	}()
}

package client_test

import (
	"packages/server/client"
	"packages/server/config"
	"testing"
)

func TestRequestCurrencyExchangeDataFromKoreaExim(t *testing.T) {
	if err := config.Load("../../settings.toml"); err != nil {
		t.Error(err)
	}

	r, err := client.RequestCurrencyExchangeDataFromKoreaExim("2025-03-01")
	if err != nil {
		t.Error(err)
	}

	for _, v := range r {
		t.Log(v.CurrencyCode, v.ExchangeRate)
	}
}

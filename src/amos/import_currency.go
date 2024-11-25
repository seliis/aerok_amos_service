package amos

import (
	"encoding/xml"
	"packages/src/config"
	"packages/src/internal/entities"
)

type _ImportCurrency struct {
	XMLName    string      `xml:"importCurrency"`
	Version    xml.Attr    `xml:"version,attr"`
	Currencies []_Currency `xml:"currency"`
}

type _Currency struct {
	XMLName      string  `xml:"currency"`
	CurrencyCode string  `xml:"currencyCode"`
	Description  string  `xml:"description"`
	ExchangeRate float64 `xml:"exchangeRate"`
	ExchangeBase uint    `xml:"exchangeBase"`
}

func NewImportCurrency(exchangeRates []entities.ExchangeRate) _ImportCurrency {
	var currencies []_Currency
	var baseRate float64

	for _, exchangeRate := range exchangeRates {
		if exchangeRate.CurrencyCode == config.Amos.ImportCurrency.BaseCurrency {
			baseRate = exchangeRate.DirectRate
			break
		}
	}

	for _, exchangeRate := range exchangeRates {
		currencies = append(currencies, _Currency{
			CurrencyCode: exchangeRate.CurrencyCode,
			Description:  exchangeRate.CurrencyName,
			ExchangeRate: exchangeRate.DirectRate / baseRate,
			ExchangeBase: 1,
		})
	}

	currencies = append(currencies, _Currency{
		CurrencyCode: config.Server.Service.BaseCurrency,
		Description:  getBaseCurrencyDescription(),
		ExchangeRate: 1 / baseRate,
		ExchangeBase: 1,
	})

	return _ImportCurrency{
		Version: xml.Attr{
			Name:  xml.Name{Local: "version"},
			Value: "0.2",
		},
		Currencies: currencies,
	}
}

func getBaseCurrencyDescription() string {
	for _, currency := range config.Amos.ImportCurrency.Currencies {
		if currency.Code == config.Server.Service.BaseCurrency {
			return currency.Name
		}
	}

	return "Unknown"
}

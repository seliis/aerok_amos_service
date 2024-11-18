package amos

import (
	"encoding/xml"
	"packages/src/config"
	"packages/src/internal/entities"
)

type importCurrency struct {
	XMLName    string     `xml:"importCurrency"`
	Version    xml.Attr   `xml:"version,attr"`
	Currencies []currency `xml:"currency"`
}

type currency struct {
	XMLName      string  `xml:"currency"`
	CurrencyCode string  `xml:"currencyCode"`
	Description  string  `xml:"description"`
	ExchangeRate float64 `xml:"exchangeRate"`
	ExchangeBase uint    `xml:"exchangeBase"`
}

func NewImportCurrency(exchangeRates []entities.ExchangeRate) importCurrency {
	var currencies []currency
	var baseRate float64

	for _, exchangeRate := range exchangeRates {
		if exchangeRate.CurrencyCode == config.Amos.ImportCurrency.BaseCurrency {
			baseRate = exchangeRate.DirectRate
			break
		}
	}

	for _, exchangeRate := range exchangeRates {
		currencies = append(currencies, currency{
			CurrencyCode: exchangeRate.CurrencyCode,
			Description:  exchangeRate.CurrencyName,
			ExchangeRate: exchangeRate.DirectRate / baseRate,
			ExchangeBase: 1,
		})
	}

	currencies = append(currencies, currency{
		CurrencyCode: config.Server.Service.BaseCurrency,
		Description:  getDescription(),
		ExchangeRate: 1 / baseRate,
		ExchangeBase: 1,
	})

	return importCurrency{
		Version: xml.Attr{
			Name:  xml.Name{Local: "version"},
			Value: "0.2",
		},
		Currencies: currencies,
	}
}

func getDescription() string {
	for _, applicable := range config.Amos.ImportCurrency.Applicables {
		if applicable.CurrencyCode == config.Server.Service.BaseCurrency {
			return applicable.CurrencyName
		}
	}

	return "Unknown"
}

package amos

import (
	"encoding/xml"
	"errors"
	"net/http"
	"packages/server/client"
	"packages/server/config"
	"packages/server/internal/entities"
)

type _ImportCurrency struct {
	XMLName    string       `xml:"importCurrency"`
	Version    xml.Attr     `xml:"version,attr"`
	Currencies []*_Currency `xml:"currency"`
}

type _Currency struct {
	XMLName      string  `xml:"currency"`
	CurrencyCode string  `xml:"currencyCode"`
	Description  string  `xml:"description"`
	ExchangeRate float64 `xml:"exchangeRate"`
	ExchangeBase uint    `xml:"exchangeBase"`
}

func NewImportCurrency(exchangeRates []*entities.ExchangeRateWithCurrency) (*_ImportCurrency, error) {
	config := config.AMOS.Services.ImportCurrency
	var amosCurrencyRate float64

	for _, exchangeRate := range exchangeRates {
		if exchangeRate.Code == config.AmosCurrencyCode {
			amosCurrencyRate = exchangeRate.Rate
			break
		}
	}

	if amosCurrencyRate == 0 {
		return nil, errors.New("amos currency rate not found")
	}

	var data []*_Currency

	for _, exchangeRate := range exchangeRates {
		data = append(data, &_Currency{
			XMLName:      "currency",
			CurrencyCode: exchangeRate.Code,
			Description:  exchangeRate.Name,
			ExchangeRate: exchangeRate.Rate / amosCurrencyRate / float64(exchangeRate.Base),
			ExchangeBase: 1,
		})
	}

	return &_ImportCurrency{
		XMLName: "importCurrency",
		Version: xml.Attr{
			Name: xml.Name{
				Local: "version",
			},
			Value: "0.2",
		},
		Currencies: data,
	}, nil
}

func (importCurrency *_ImportCurrency) Push() error {
	config := config.AMOS.Services.ImportCurrency

	r, err := client.GetAmosRequest().SetBody(importCurrency).Post(config.EndPoint)
	if err != nil {
		return err
	}

	if r.StatusCode() != http.StatusOK {
		return errors.New("failed to push data")
	}

	return nil
}

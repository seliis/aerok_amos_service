package repositories

import (
	"packages/src/amos"
	"packages/src/config"
	"packages/src/infra/client"
	"packages/src/infra/database"
	"packages/src/internal/entities"
)

type CurrencyRepository struct{}

func NewCurrencyRepository() *CurrencyRepository {
	return &CurrencyRepository{}
}

func (r *CurrencyRepository) GetPrimitiveExchangeRates(date string) ([]entities.ExchangeRatePrimitive, error) {
	var primitives []entities.ExchangeRatePrimitive

	client := client.NewWithRetry("https://www.koreaexim.go.kr")

	_, err := client.R().
		SetQueryParams(map[string]string{
			"authkey":    config.Server.Service.KoreaEximAuthKey,
			"searchdate": date,
			"data":       "AP01",
		}).
		SetResult(&primitives).
		Get("/site/program/financial/exchangeJSON")

	if err != nil {
		return nil, err
	}

	return primitives, nil
}

func (r *CurrencyRepository) UpdateAmos(exchangeRates []entities.ExchangeRate) error {
	client := client.New(config.Amos.BaseUrl)

	_, err := client.R().
		SetHeader("Content-Type", "application/xml").
		SetBasicAuth(config.Amos.WebserviceId, config.Amos.WebservicePassword).
		SetBody(amos.NewImportCurrency(exchangeRates)).
		Post(config.Amos.ImportCurrency.Endpoint)

	if err != nil {
		return err
	}

	return nil
}

func (r *CurrencyRepository) SaveExchangeRate(data []entities.ExchangeRate) error {
	statement := `
			insert or replace
			into exchange_rates (
				currency_code,
				currency_name,
				date,
				direct_rate
			) values (
			 	?, ?, ?, ?
			) on conflict(
				currency_code,
				date
			) do update set
				direct_rate = excluded.direct_rate
		`

	for _, e := range data {
		if _, err := database.Client.Exec(statement, e.CurrencyCode, e.CurrencyName, e.Date, e.DirectRate); err != nil {
			return err
		}
	}

	return nil
}

func (r *CurrencyRepository) GetExchangeRate(currencyCode, date string) (*entities.ExchangeRate, error) {
	var exchangeRate entities.ExchangeRate

	if err := database.Client.QueryRow("select currency_code, currency_name, date, direct_rate from exchange_rates where currency_code = ? and date = ?", currencyCode, date).Scan(&exchangeRate.CurrencyCode, &exchangeRate.CurrencyName, &exchangeRate.Date, &exchangeRate.DirectRate); err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func (r *CurrencyRepository) IsExists(currencyCode, date string) (bool, error) {
	var count int

	if err := database.Client.QueryRow("select count(*) from exchange_rates where currency_code = ? and date = ?", currencyCode, date).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

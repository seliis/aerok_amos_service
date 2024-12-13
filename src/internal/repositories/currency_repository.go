package repositories

import (
	"packages/src/config"
	"packages/src/internal/entities"
	"packages/src/pkg/client"
	"packages/src/pkg/database"
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

func (r *CurrencyRepository) SaveExchangeRate(data []entities.ExchangeRate) error {
	statement := `
			insert or replace
			into exchange_rates (
				currency_code,
				currency_name,
				currency_date,
				direct_rate
			) values (
			 	?, ?, ?, ?
			) on conflict(
				currency_code,
				currency_date
			) do update set
				direct_rate = excluded.direct_rate
		`

	for _, e := range data {
		if _, err := database.Client.Exec(statement, e.CurrencyCode, e.CurrencyName, e.CurrencyDate, e.DirectRate); err != nil {
			return err
		}
	}

	return nil
}

func (r *CurrencyRepository) GetExchangeRate(currencyCode, date string) (*entities.ExchangeRate, error) {
	var exchangeRate entities.ExchangeRate

	if err := database.Client.QueryRow("select * from exchange_rates where currency_code = ? and currency_date = ?", currencyCode, date).Scan(&exchangeRate.CurrencyCode, &exchangeRate.CurrencyName, &exchangeRate.CurrencyDate, &exchangeRate.DirectRate); err != nil {
		return nil, err
	}

	return &exchangeRate, nil
}

func (r *CurrencyRepository) IsExists(currencyCode, date string) (bool, error) {
	var count int

	if err := database.Client.QueryRow("select count(*) from exchange_rates where currency_code = ? and currency_date = ?", currencyCode, date).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

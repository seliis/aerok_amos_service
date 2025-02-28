package entities

type ExchangeRate struct {
	ID         string  `json:"id"`
	Date       string  `json:"date"`
	Rate       float64 `json:"rate"`
	CurrencyID string  `json:"currency_id"`
}

type ExchangeRateWithCurrency struct {
	ExchangeRateID string  `json:"exchange_rate_id"`
	CurrencyID     string  `json:"currency_id"`
	Code           string  `json:"code"`
	Name           string  `json:"name"`
	Base           int     `json:"base"`
	Date           string  `json:"date"`
	Rate           float64 `json:"rate"`
}

package entities

type ExchangeRate struct {
	ID           string  `json:"id"`
	CurrencyCode string  `json:"currency_code"`
	Date         string  `json:"date"`
	Rate         float64 `json:"rate"`
}

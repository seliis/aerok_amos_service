package entities

type ExchangeRate struct {
	ID           string  `json:"id"`
	CurrencyCode string  `json:"currency_code"`
	Date         string  `json:"date"`
	Rate         float64 `json:"rate"`
}

type ExchangeRateWithCurrency struct {
	ID   string  `json:"id"`
	Code string  `json:"code"`
	Name string  `json:"name"`
	Base int     `json:"base"`
	Date string  `json:"date"`
	Rate float64 `json:"rate"`
}

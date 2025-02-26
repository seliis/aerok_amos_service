package entities

import "time"

type ExchangeRate struct {
	ID         string    `json:"id"`
	Date       time.Time `json:"date"`
	Rate       float64   `json:"rate"`
	CurrencyID string    `json:"currency_id"`
}

package entities

type CurrencyCode struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Base int    `json:"base"`
}

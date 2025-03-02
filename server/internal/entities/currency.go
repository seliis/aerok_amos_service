package entities

type Currency struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Base int    `json:"base"`
}

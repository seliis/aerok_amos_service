package repositories

import (
	"packages/src/amos"
	"packages/src/config"
	"packages/src/internal/entities"
	"packages/src/pkg/client"
)

type AmosRepository struct{}

func NewAmosRepository() *AmosRepository {
	return &AmosRepository{}
}

func (r *AmosRepository) UpdateCurrency(authorization string, exchangeRates []entities.ExchangeRate) error {
	client := client.New(config.Amos.BaseUrl)

	_, err := client.R().
		SetHeader("Authorization", authorization).
		SetHeader("Content-Type", "application/xml").
		SetBody(amos.NewImportCurrency(exchangeRates)).
		Post(config.Amos.ImportCurrency.Endpoint)

	if err != nil {
		return err
	}

	return nil
}

package services

import (
	"context"
	"packages/server/amos"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
)

type AmosService struct {
	_ExchangeRateRepository *repositories.ExchangeRateRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		_ExchangeRateRepository: repositories.NewExchangeRateRepository(),
	}
}

func (s *AmosService) ImportCurrency(context context.Context, exchangeRates []*entities.ExchangeRateWithCurrency) error {
	data, err := amos.NewImportCurrency(exchangeRates)
	if err != nil {
		return err
	}

	if err := data.Push(); err != nil {
		return err
	}

	return nil
}

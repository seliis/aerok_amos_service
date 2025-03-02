package services

import (
	"context"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
)

type CurrencyService struct {
	_CurrencyRepository *repositories.CurrencyRepository
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		_CurrencyRepository: repositories.NewCurrencyRepository(),
	}
}

func (s *CurrencyService) Create(context context.Context, entity *entities.Currency) (*entities.Currency, error) {
	return s._CurrencyRepository.Create(context, entity)
}

func (s *CurrencyService) Read(context context.Context, currencyCode string) (*entities.Currency, error) {
	return s._CurrencyRepository.Read(context, currencyCode)
}

func (s *CurrencyService) Update(context context.Context, entity *entities.Currency) (*entities.Currency, error) {
	return s._CurrencyRepository.Update(context, entity)
}

func (s *CurrencyService) Delete(context context.Context, currencyCode string) (*entities.Currency, error) {
	return s._CurrencyRepository.Delete(context, currencyCode)
}

func (s *CurrencyService) All(context context.Context) ([]*entities.Currency, error) {
	return s._CurrencyRepository.All(context)
}

func (s *CurrencyService) Import(context context.Context, arr []*entities.Currency) ([]*entities.Currency, error) {
	var results []*entities.Currency

	for _, entity := range arr {
		result, err := s.Create(context, entity)
		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

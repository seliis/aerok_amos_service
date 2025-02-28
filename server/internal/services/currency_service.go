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

func (s *CurrencyService) Read(context context.Context, id string) (*entities.Currency, error) {
	return s._CurrencyRepository.Read(context, id)
}

func (s *CurrencyService) Update(context context.Context, entity *entities.Currency) (*entities.Currency, error) {
	return s._CurrencyRepository.Update(context, entity)
}

func (s *CurrencyService) Delete(context context.Context, id string) (*entities.Currency, error) {
	return s._CurrencyRepository.Delete(context, id)
}

func (s *CurrencyService) All(context context.Context) ([]*entities.Currency, error) {
	return s._CurrencyRepository.All(context)
}

func (s *CurrencyService) Import(context context.Context, entities []*entities.Currency) ([]*entities.Currency, error) {
	return s._CurrencyRepository.Import(context, entities)
}

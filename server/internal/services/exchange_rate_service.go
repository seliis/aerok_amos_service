package services

import (
	"context"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
)

type ExchangeRateService struct {
	_ExchangeRateRepository *repositories.ExchangeRateRepository
}

func NewExchangeRateService() *ExchangeRateService {
	return &ExchangeRateService{
		_ExchangeRateRepository: repositories.NewExchangeRateRepository(),
	}
}

func (s *ExchangeRateService) Create(context context.Context, entity *entities.ExchangeRate) (*entities.ExchangeRate, error) {
	return s._ExchangeRateRepository.Create(context, entity)
}

func (s *ExchangeRateService) Read(context context.Context, id string) (*entities.ExchangeRate, error) {
	return s._ExchangeRateRepository.Read(context, id)
}

func (s *ExchangeRateService) Update(context context.Context, entity *entities.ExchangeRate) (*entities.ExchangeRate, error) {
	return s._ExchangeRateRepository.Update(context, entity)
}

func (s *ExchangeRateService) Delete(context context.Context, id string) (*entities.ExchangeRate, error) {
	return s._ExchangeRateRepository.Delete(context, id)
}

func (s *ExchangeRateService) All(context context.Context) ([]*entities.ExchangeRate, error) {
	return s._ExchangeRateRepository.All(context)
}

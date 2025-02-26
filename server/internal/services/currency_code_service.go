package services

import (
	"context"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
)

type CurrencyCodeService struct {
	_CurrencyCodeRepository *repositories.CurrencyCodeRepository
}

func NewCurrencyCodeService() *CurrencyCodeService {
	return &CurrencyCodeService{
		_CurrencyCodeRepository: repositories.NewCurrencyCodeRepository(),
	}
}

func (s *CurrencyCodeService) Create(context context.Context, entity *entities.CurrencyCode) (*entities.CurrencyCode, error) {
	return s._CurrencyCodeRepository.Create(context, entity)
}

func (s *CurrencyCodeService) Read(context context.Context, id string) (*entities.CurrencyCode, error) {
	return s._CurrencyCodeRepository.Read(context, id)
}

func (s *CurrencyCodeService) Update(context context.Context, entity *entities.CurrencyCode) (*entities.CurrencyCode, error) {
	return s._CurrencyCodeRepository.Update(context, entity)
}

func (s *CurrencyCodeService) Delete(context context.Context, id string) (*entities.CurrencyCode, error) {
	return s._CurrencyCodeRepository.Delete(context, id)
}

func (s *CurrencyCodeService) All(context context.Context) ([]*entities.CurrencyCode, error) {
	return s._CurrencyCodeRepository.All(context)
}

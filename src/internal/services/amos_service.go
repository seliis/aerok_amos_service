package services

import (
	"encoding/base64"
	"packages/src/internal/dto"
	"packages/src/internal/entities"
	"packages/src/internal/repositories"

	"github.com/gin-gonic/gin"
)

type AmosService struct {
	amosRepository     *repositories.AmosRepository
	currencyRepository *repositories.CurrencyRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		amosRepository:     repositories.NewAmosRepository(),
		currencyRepository: repositories.NewCurrencyRepository(),
	}
}

func (s *AmosService) GetBasicAuth(context *gin.Context, id, password string) string {
	return base64.StdEncoding.EncodeToString([]byte(id + ":" + password))
}

func (s *AmosService) UpdateCurrency(context *gin.Context, request *dto.UpdateAmosCurrencyRequest, data []entities.ExchangeRate) error {
	if err := s.amosRepository.UpdateCurrency(request.Authorization, data); err != nil {
		return err
	}

	return nil
}

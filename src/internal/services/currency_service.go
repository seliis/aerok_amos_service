package services

import (
	"packages/src/config"
	"packages/src/internal/dto"
	"packages/src/internal/entities"
	"packages/src/internal/repositories"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CurrencyService struct {
	currencyRepository *repositories.CurrencyRepository
}

func NewCurrencyService() *CurrencyService {
	return &CurrencyService{
		currencyRepository: repositories.NewCurrencyRepository(),
	}
}

func (s *CurrencyService) GetExchangeRates(context *gin.Context, request *dto.GetExchangeRatesRequest) ([]entities.ExchangeRate, error) {
	primitives, err := s.currencyRepository.GetPrimitiveExchangeRates(request.Date)
	if err != nil {
		return nil, err
	}

	applicables := config.Amos.ImportCurrency.Applicables
	var data []entities.ExchangeRate

	for _, primitive := range primitives {
		if primitive.CurrencyCode == config.Server.Service.BaseCurrency {
			continue
		}

		rate, err := strconv.ParseFloat(strings.ReplaceAll(primitive.ExchangeRate, ",", ""), 64)
		if err != nil {
			return nil, err
		}

		if strings.Contains(primitive.CurrencyCode, "(100)") {
			primitive.CurrencyCode = strings.ReplaceAll(primitive.CurrencyCode, "(100)", "")
			rate = rate / 100
		}

		for _, applicable := range applicables {
			if applicable.CurrencyCode == primitive.CurrencyCode {

				data = append(data, entities.ExchangeRate{
					CurrencyCode: primitive.CurrencyCode,
					CurrencyName: applicable.CurrencyName,
					Date:         request.Date,
					DirectRate:   rate,
				})
				break
			}
		}

	}

	if err := s.currencyRepository.SaveExchangeRate(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *CurrencyService) UpdateAmos(context *gin.Context, request *dto.GetExchangeRatesRequest) error {
	data, err := s.GetExchangeRates(context, request)
	if err != nil {
		return err
	}

	if err := s.currencyRepository.UpdateAmos(data); err != nil {
		return err
	}

	return nil
}

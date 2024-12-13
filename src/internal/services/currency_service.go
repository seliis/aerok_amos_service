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

	currencies := config.Amos.ImportCurrency.Currencies
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

		for _, currency := range currencies {
			if currency.Code == primitive.CurrencyCode {

				data = append(data, entities.ExchangeRate{
					CurrencyCode: primitive.CurrencyCode,
					CurrencyName: currency.Name,
					CurrencyDate: request.Date,
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

func (s *CurrencyService) GetExchangeRate(context *gin.Context, request *dto.GetExchangeRateRequest) (*entities.ExchangeRate, error) {
	exists, err := s.currencyRepository.IsExists(request.CurrencyCode, request.Date)
	if err != nil {
		return nil, err
	}

	if !exists {
		_, err := s.GetExchangeRates(context, &dto.GetExchangeRatesRequest{
			Date: request.Date,
		})
		if err != nil {
			return nil, err
		}
	}

	return s.currencyRepository.GetExchangeRate(request.CurrencyCode, request.Date)
}

func (s *CurrencyService) GetCurrencies(context *gin.Context) []map[string]string {
	var data []map[string]string

	for _, currency := range config.Amos.ImportCurrency.Currencies {
		if currency.Code != config.Server.Service.BaseCurrency {
			data = append(data, map[string]string{
				"code": currency.Code,
				"name": currency.Name,
			})
		}
	}

	return data
}

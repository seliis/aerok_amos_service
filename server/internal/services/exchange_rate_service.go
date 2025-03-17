package services

import (
	"context"
	"errors"
	"fmt"
	"packages/server/client"
	"packages/server/domain"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
	"strconv"
	"strings"
	"time"
)

type ExchangeRateService struct {
	_ExchangeRateRepository *repositories.ExchangeRateRepository
	_CurrencyRepository     *repositories.CurrencyRepository
}

func NewExchangeRateService() *ExchangeRateService {
	return &ExchangeRateService{
		_ExchangeRateRepository: repositories.NewExchangeRateRepository(),
		_CurrencyRepository:     repositories.NewCurrencyRepository(),
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

func (s *ExchangeRateService) UpdateExchangeRates(context context.Context, date string) error {
	primitives, err := client.RequestCurrencyExchangeDataFromKoreaExim(date)
	if err != nil {
		return err
	}

	if len(primitives) == 0 {
		return domain.ErrKoreaEximExchangeRateNotFound
	}

	currencies, err := s._CurrencyRepository.All(context)
	if err != nil {
		return err
	}

	for _, currency := range currencies {
		for _, primitive := range primitives {
			if strings.Contains(primitive.CurrencyCode, currency.Code) {
				exchangeRate, err := strconv.ParseFloat(strings.ReplaceAll(primitive.ExchangeRate, ",", ""), 64)
				if err != nil {
					return err
				}

				entity := &entities.ExchangeRate{
					Date:         date,
					Rate:         exchangeRate,
					CurrencyCode: currency.Code,
				}

				if err = s._ExchangeRateRepository.UpsertExchangeRate(context, entity); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (s *ExchangeRateService) UpdateAnnualExchangeRates(context context.Context, year string) error {
	today := time.Now()

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return err
	}

	if yearInt < 1991 || yearInt > today.Year() {
		return errors.New("year must be between 1991 and current year")
	}

	start, err := time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))
	if err != nil {
		return err
	}

	end, err := time.Parse("2006-01-02", fmt.Sprintf("%s-12-31", year))
	if err != nil {
		return err
	}

	if today.Year() == start.Year() && today.Before(end) {
		end = today
	}

	for date := start; !date.After(end); date = date.AddDate(0, 0, 1) {
		err := s.UpdateExchangeRates(context, date.Format("2006-01-02"))
		if err != nil && err != domain.ErrKoreaEximExchangeRateNotFound {
			return err
		}
	}

	return nil
}

func (s *ExchangeRateService) GetExchangeRates(context context.Context, date string) ([]*entities.ExchangeRateWithCurrency, error) {
	currencies, err := s._CurrencyRepository.All(context)
	if err != nil {
		return nil, err
	}

	arr, err := s._ExchangeRateRepository.ListByDate(context, date)
	if err != nil {
		return nil, err
	}

	if len(arr) != len(currencies) {
		err := s.UpdateExchangeRates(context, date)
		if err != nil {
			return nil, err
		}
	}

	exchangeRates, err := s._ExchangeRateRepository.GetExchangeRates(context, date)
	if err != nil {
		return nil, err
	}

	return exchangeRates, nil
}

func (s *ExchangeRateService) GetLatestExchangeRate(context context.Context, code string, date string) (*entities.ExchangeRateWithCurrency, error) {
	isExist, err := s._ExchangeRateRepository.IsExist(context, code, date)
	if err != nil {
		return nil, err
	}

	if !isExist {
		err := s.UpdateExchangeRates(context, date)
		if err != nil && err != domain.ErrKoreaEximExchangeRateNotFound {
			return nil, err
		}
	}

	exchangeRate, err := s._ExchangeRateRepository.GetLatestExchangeRate(context, code, date)
	if err != nil {
		return nil, err
	}

	return exchangeRate, nil
}

func (s *ExchangeRateService) GetAnnualExchangeRates(context context.Context, code, year string) ([]*entities.ExchangeRateWithCurrency, error) {
	arr, err := s._ExchangeRateRepository.GetAnnualExchangeRates(context, code, year)
	if err != nil {
		return nil, err
	}

	return arr, nil
}

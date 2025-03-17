package services

import (
	"context"
	"encoding/base64"
	"packages/server/amos"
	"packages/server/config"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
	"strings"
	"time"
)

type AmosService struct {
	_FlightScheduleRepository *repositories.FlightScheduleRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		_FlightScheduleRepository: repositories.NewFlightScheduleRepository(),
	}
}

func (s *AmosService) Authorize(header string) (string, bool) {
	if !strings.HasPrefix(header, "Basic ") {
		return "", false
	}

	token := strings.TrimPrefix(header, "Basic ")

	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", false
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return "", false
	}

	if credentials[0] != config.AMOS.Auth.ID {
		return "", false
	}

	if credentials[1] != config.AMOS.Auth.Password {
		return "", false
	}

	return token, true
}

func (s *AmosService) ImportCurrency(context context.Context, token string, exchangeRates []*entities.ExchangeRateWithCurrency) error {
	data, err := amos.NewImportCurrency(exchangeRates)
	if err != nil {
		return err
	}

	if err := data.Push(token); err != nil {
		return err
	}

	return nil
}

func (s *AmosService) TransferFutureFlights(context context.Context, token string) error {
	date := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	flightSchedules, err := s._FlightScheduleRepository.GetFlightSchedulesFromDate(context, date)
	if err != nil {
		return err
	}

	data := amos.NewFutureFlights(flightSchedules)
	if err := data.Push(token); err != nil {
		return err
	}

	return nil
}

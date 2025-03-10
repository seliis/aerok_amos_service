package services

import (
	"context"
	"encoding/base64"
	"packages/server/amos"
	"packages/server/config"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
	"strings"
)

type AmosService struct {
	_FlightScheduleRepository *repositories.FlightScheduleRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		_FlightScheduleRepository: repositories.NewFlightScheduleRepository(),
	}
}

func (s *AmosService) Authorize(header string) bool {
	if !strings.HasPrefix(header, "Basic ") {
		return false
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(header, "Basic "))
	if err != nil {
		return false
	}

	credentials := strings.SplitN(string(decoded), ":", 2)
	if len(credentials) != 2 {
		return false
	}

	if credentials[0] != config.AMOS.Auth.ID {
		return false
	}

	if credentials[1] != config.AMOS.Auth.Password {
		return false
	}

	return true
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

func (s *AmosService) TransferFutureFlights(context context.Context, date string) error {
	flightSchedules, err := s._FlightScheduleRepository.GetFlightSchedulesFromDate(context, date)
	if err != nil {
		return err
	}

	data := amos.NewFutureFlights(flightSchedules)
	if err := data.Push(); err != nil {
		return err
	}

	return nil
}

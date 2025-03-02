package services

import (
	"context"
	"packages/server/amos"
	"packages/server/config"
	"packages/server/internal/entities"
	"packages/server/internal/repositories"
)

type AmosService struct {
	_FlightScheduleRepository *repositories.FlightScheduleRepository
}

func NewAmosService() *AmosService {
	return &AmosService{
		_FlightScheduleRepository: repositories.NewFlightScheduleRepository(),
	}
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

func (s *AmosService) TransferFutureFlights(context context.Context, fromDate string) error {
	daysAhead := config.AMOS.Services.TransferFutureFlights.DaysAhead

	flightSchedules, err := s._FlightScheduleRepository.GetFlightSchedulesByPeriod(context, fromDate, daysAhead)
	if err != nil {
		return err
	}

	data := amos.NewFutureFlights(flightSchedules)
	if err := data.Push(); err != nil {
		return err
	}

	return nil
}

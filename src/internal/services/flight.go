package services

import "packages/src/internal/repositories"

type FlightService struct {
	flightRepository *repositories.FlightRepository
}

func NewFlightService() *FlightService {
	return &FlightService{flightRepository: repositories.NewFlightRepository()}
}

func (s *FlightService) UpdateAmos() error {
	return s.flightRepository.UpdateAmos("admin", "8Z7Plc3p!!")
}

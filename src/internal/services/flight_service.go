package services

import (
	"packages/src/internal/repositories"
)

type FlightService struct {
	flightRepository *repositories.FlightRepository
}

func NewFlightService() *FlightService {
	return &FlightService{flightRepository: repositories.NewFlightRepository()}
}

package services

import (
	"packages/src/amos"
	"packages/src/internal/dto"
	"packages/src/internal/repositories"

	"github.com/gin-gonic/gin"
)

type FlightService struct {
	flightRepository *repositories.FlightRepository
}

func NewFlightService() *FlightService {
	return &FlightService{flightRepository: repositories.NewFlightRepository()}
}

func (s *FlightService) UpdateAmos(context *gin.Context, request *dto.UpdateAmosFutureFlightsRequest) error {
	if err := amos.CheckWebServiceAuth(request.WebServiceAuth); err != nil {
		return err
	}

	return s.flightRepository.UpdateAmos(request.WebServiceAuth.Id, request.WebServiceAuth.Password)
}

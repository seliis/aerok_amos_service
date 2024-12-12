package handlers

import (
	"packages/src/internal/services"
)

type FlightHandler struct {
	flightService *services.FlightService
}

func NewFlightHandler() *FlightHandler {
	return &FlightHandler{flightService: services.NewFlightService()}
}

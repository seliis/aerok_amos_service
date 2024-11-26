package handlers

import (
	"net/http"
	"packages/src/internal/dto"
	"packages/src/internal/services"

	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	flightService *services.FlightService
}

func NewFlightHandler() *FlightHandler {
	return &FlightHandler{flightService: services.NewFlightService()}
}

func (h *FlightHandler) UpdateAmos(c *gin.Context) {
	request, err := dto.NewUpdateAmosFutureFlightsRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.flightService.UpdateAmos(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(nil))
}

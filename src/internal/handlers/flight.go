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
	if err := h.flightService.UpdateAmos(); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
	} else {
		c.JSON(http.StatusOK, dto.NewSuccessResponse(nil))
	}
}

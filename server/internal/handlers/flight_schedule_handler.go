package handlers

import (
	"net/http"
	"packages/server/internal/requests"
	"packages/server/internal/responses"
	"packages/server/internal/services"

	"github.com/gin-gonic/gin"
)

type FlightScheduleHandler struct {
	_FlightScheduleService *services.FlightScheduleService
}

func NewFlightScheduleHandler() *FlightScheduleHandler {
	return &FlightScheduleHandler{
		_FlightScheduleService: services.NewFlightScheduleService(),
	}
}

func (h *FlightScheduleHandler) UpdateFlightSchedules(context *gin.Context) {
	workbook, err := requests.GetWorkbook(context, "flight_schedule")
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	if err := h._FlightScheduleService.UpdateFlightSchedulesFromWorkbook(context, workbook); err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(nil))
}

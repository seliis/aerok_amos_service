package handlers

import (
	"errors"
	"net/http"
	"packages/server/internal/requests"
	"packages/server/internal/responses"
	"packages/server/internal/services"

	"github.com/gin-gonic/gin"
)

type AmosHandler struct {
	_AmosService         *services.AmosService
	_ExchangeRateService *services.ExchangeRateService
}

func NewAmosHandler() *AmosHandler {
	return &AmosHandler{
		_AmosService:         services.NewAmosService(),
		_ExchangeRateService: services.NewExchangeRateService(),
	}
}

func (h *AmosHandler) ImportCurrency(context *gin.Context) {
	date, err := requests.GetDate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	exchangeRates, err := h._ExchangeRateService.GetExchangeRates(context, date)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	if err := h._AmosService.ImportCurrency(context, exchangeRates); err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(nil))
}

func (h *AmosHandler) TransferFutureFlights(context *gin.Context) {
	fromDate, isOk := context.GetQuery("from")
	if !isOk {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(errors.New("from_date is required")))
		return
	}

	if err := h._AmosService.TransferFutureFlights(context, fromDate); err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(nil))
}

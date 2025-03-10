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

func (h *AmosHandler) Authorize(context *gin.Context) {
	if !h._AmosService.Authorize(requests.GetAuthorization(context)) {
		context.JSON(http.StatusUnauthorized, responses.NewErrorResponse(errors.New("Unauthorized")))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse("Authorized"))
}

func (h *AmosHandler) ImportCurrency(context *gin.Context) {
	if !h._AmosService.Authorize(requests.GetAuthorization(context)) {
		context.JSON(http.StatusUnauthorized, responses.NewErrorResponse(errors.New("Unauthorized")))
		return
	}

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
	if !h._AmosService.Authorize(requests.GetAuthorization(context)) {
		context.JSON(http.StatusUnauthorized, responses.NewErrorResponse(errors.New("Unauthorized")))
		return
	}

	date, err := requests.GetDate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	if err := h._AmosService.TransferFutureFlights(context, date); err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(nil))
}

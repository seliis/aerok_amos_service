package handlers

import (
	"errors"
	"net/http"
	"packages/server/internal/entities"
	"packages/server/internal/requests"
	"packages/server/internal/responses"
	"packages/server/internal/services"

	"github.com/gin-gonic/gin"
)

type ExchangeRateHandler struct {
	_ExchangeRateService *services.ExchangeRateService
}

func NewExchangeRateHandler() *ExchangeRateHandler {
	return &ExchangeRateHandler{
		_ExchangeRateService: services.NewExchangeRateService(),
	}
}

func (h *ExchangeRateHandler) Create(context *gin.Context) {
	var entity *entities.ExchangeRate

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._ExchangeRateService.Create(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusCreated, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) Read(context *gin.Context) {
	result, err := h._ExchangeRateService.Read(context, requests.GetID(context))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) Update(context *gin.Context) {
	var entity *entities.ExchangeRate

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._ExchangeRateService.Update(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) Delete(context *gin.Context) {
	result, err := h._ExchangeRateService.Delete(context, requests.GetID(context))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) All(context *gin.Context) {
	result, err := h._ExchangeRateService.All(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) UpdateExchangeRates(context *gin.Context) {
	date, err := requests.GetDate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	if err := h._ExchangeRateService.UpdateExchangeRates(context, date); err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(nil))
}

func (h *ExchangeRateHandler) GetExchangeRates(context *gin.Context) {
	date, err := requests.GetDate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._ExchangeRateService.GetExchangeRates(context, date)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *ExchangeRateHandler) GetExchangeRate(context *gin.Context) {
	currencyCode, isOk := context.GetQuery("code")
	if !isOk {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(errors.New("code is required")))
		return
	}

	date, err := requests.GetDate(context)
	if err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._ExchangeRateService.GetExchangeRate(context, currencyCode, date)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

package handlers

import (
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

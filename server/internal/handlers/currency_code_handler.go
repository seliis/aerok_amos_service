package handlers

import (
	"net/http"
	"packages/server/internal/entities"
	"packages/server/internal/requests"
	"packages/server/internal/responses"
	"packages/server/internal/services"

	"github.com/gin-gonic/gin"
)

type CurrencyCodeHandler struct {
	_CurrencyCodeService *services.CurrencyCodeService
}

func NewCurrencyCodeHandler() *CurrencyCodeHandler {
	return &CurrencyCodeHandler{
		_CurrencyCodeService: services.NewCurrencyCodeService(),
	}
}

func (h *CurrencyCodeHandler) Create(context *gin.Context) {
	var entity *entities.CurrencyCode

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._CurrencyCodeService.Create(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusCreated, responses.NewSuccessResponse(result))
}

func (h *CurrencyCodeHandler) Read(context *gin.Context) {
	result, err := h._CurrencyCodeService.Read(context, requests.GetID(context))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyCodeHandler) Update(context *gin.Context) {
	var entity *entities.CurrencyCode

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._CurrencyCodeService.Update(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyCodeHandler) Delete(context *gin.Context) {
	result, err := h._CurrencyCodeService.Delete(context, requests.GetID(context))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyCodeHandler) All(context *gin.Context) {
	result, err := h._CurrencyCodeService.All(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

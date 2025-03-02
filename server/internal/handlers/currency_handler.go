package handlers

import (
	"net/http"
	"packages/server/internal/entities"
	"packages/server/internal/responses"
	"packages/server/internal/services"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	_CurrencyService *services.CurrencyService
}

func NewCurrencyHandler() *CurrencyHandler {
	return &CurrencyHandler{
		_CurrencyService: services.NewCurrencyService(),
	}
}

func (h *CurrencyHandler) Create(context *gin.Context) {
	var entity *entities.Currency

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._CurrencyService.Create(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusCreated, responses.NewSuccessResponse(result))
}

func (h *CurrencyHandler) Read(context *gin.Context) {
	result, err := h._CurrencyService.Read(context, context.Param("currency_code"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyHandler) Update(context *gin.Context) {
	var entity *entities.Currency

	if err := context.ShouldBindBodyWithJSON(&entity); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._CurrencyService.Update(context, entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyHandler) Delete(context *gin.Context) {
	result, err := h._CurrencyService.Delete(context, context.Param("currency_code"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyHandler) All(context *gin.Context) {
	result, err := h._CurrencyService.All(context)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusOK, responses.NewSuccessResponse(result))
}

func (h *CurrencyHandler) Import(context *gin.Context) {
	var entities []*entities.Currency

	if err := context.ShouldBindBodyWithJSON(&entities); err != nil {
		context.JSON(http.StatusBadRequest, responses.NewErrorResponse(err))
		return
	}

	result, err := h._CurrencyService.Import(context, entities)
	if err != nil {
		context.JSON(http.StatusInternalServerError, responses.NewErrorResponse(err))
		return
	}

	context.JSON(http.StatusCreated, responses.NewSuccessResponse(result))
}

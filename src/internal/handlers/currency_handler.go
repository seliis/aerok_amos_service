package handlers

import (
	"net/http"
	"packages/src/internal/dto"
	"packages/src/internal/services"

	"github.com/gin-gonic/gin"
)

type CurrencyHandler struct {
	currencyService *services.CurrencyService
}

func NewCurrencyHandler() *CurrencyHandler {
	return &CurrencyHandler{
		currencyService: services.NewCurrencyService(),
	}
}

func (h *CurrencyHandler) GetExchangeRates(c *gin.Context) {
	request, err := dto.NewGetExchangeRatesRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if data, err := h.currencyService.GetExchangeRates(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	} else {
		c.JSON(http.StatusOK, dto.NewSuccessResponse(data))
	}
}

func (h *CurrencyHandler) GetExchangeRate(c *gin.Context) {
	request, err := dto.NewGetExchangeRateRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if data, err := h.currencyService.GetExchangeRate(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	} else {
		c.JSON(http.StatusOK, dto.NewSuccessResponse(data))
	}
}

func (h *CurrencyHandler) UpdateAmos(c *gin.Context) {
	request, err := dto.NewGetExchangeRatesRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	if err := h.currencyService.UpdateAmos(c, request); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(nil))
}

func (h *CurrencyHandler) GetCurrencies(c *gin.Context) {
	data := h.currencyService.GetCurrencies(c)
	c.JSON(http.StatusOK, dto.NewSuccessResponse(data))
}

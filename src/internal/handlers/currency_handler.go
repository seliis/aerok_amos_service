package handlers

import (
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
		c.JSON(400, dto.NewErrorResponse(err))
		return
	}

	if data, err := h.currencyService.GetExchangeRates(c, request); err != nil {
		c.JSON(500, dto.NewErrorResponse(err))
		return
	} else {
		c.JSON(200, dto.NewSuccessResponse(data))
	}
}

func (h *CurrencyHandler) UpdateAmos(c *gin.Context) {
	request, err := dto.NewGetExchangeRatesRequest(c)
	if err != nil {
		c.JSON(400, dto.NewErrorResponse(err))
		return
	}

	if err := h.currencyService.UpdateAmos(c, request); err != nil {
		c.JSON(500, dto.NewErrorResponse(err))
		return
	}

	c.JSON(200, dto.NewSuccessResponse(nil))
}

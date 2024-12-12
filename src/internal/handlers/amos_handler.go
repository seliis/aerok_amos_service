package handlers

import (
	"net/http"
	"packages/src/internal/dto"
	"packages/src/internal/services"

	"github.com/gin-gonic/gin"
)

type AmosHandler struct {
	id              string
	password        string
	amosService     *services.AmosService
	currencyService *services.CurrencyService
}

func NewAmosHandler(id, password string) *AmosHandler {
	return &AmosHandler{
		id:              id,
		password:        password,
		amosService:     services.NewAmosService(),
		currencyService: services.NewCurrencyService(),
	}
}

func (h *AmosHandler) GetBasicAuth(c *gin.Context) {
	c.JSON(http.StatusOK, dto.NewSuccessResponse(
		h.amosService.GetBasicAuth(c, h.id, h.password),
	))
}

func (h *AmosHandler) UpdateCurrency(c *gin.Context) {
	request, err := dto.NewUpdateAmosCurrencyRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err))
		return
	}

	data, err := h.currencyService.GetExchangeRates(c, &dto.GetExchangeRatesRequest{
		Date: request.Date,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	if err := h.amosService.UpdateCurrency(c, request, data); err != nil {
		c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, dto.NewSuccessResponse(nil))
}

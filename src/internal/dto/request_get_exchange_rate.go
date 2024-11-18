package dto

import "github.com/gin-gonic/gin"

type GetExchangeRateRequest struct {
	CurrencyCode string `form:"currencyCode" binding:"required"`
	Date         string `form:"date" binding:"required"`
}

func NewGetExchangeRateRequest(c *gin.Context) (*GetExchangeRateRequest, error) {
	request := &GetExchangeRateRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		return nil, err
	}

	return request, nil
}

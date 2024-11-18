package dto

import "github.com/gin-gonic/gin"

type GetExchangeRatesRequest struct {
	Date string `form:"date" binding:"required"`
}

func NewGetExchangeRatesRequest(c *gin.Context) (*GetExchangeRatesRequest, error) {
	request := &GetExchangeRatesRequest{}

	if err := c.ShouldBindQuery(request); err != nil {
		return nil, err
	}

	return request, nil
}

package dto

import (
	"github.com/gin-gonic/gin"
)

type UpdateAmosCurrencyRequest struct {
	Authorization string
	Date          string
}

func NewUpdateAmosCurrencyRequest(c *gin.Context) (*UpdateAmosCurrencyRequest, error) {
	request := &UpdateAmosCurrencyRequest{}

	request.Authorization = c.GetHeader("Authorization")
	request.Date = c.Param("date")

	return request, nil
}

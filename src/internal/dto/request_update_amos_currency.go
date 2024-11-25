package dto

import (
	"packages/src/amos"

	"github.com/gin-gonic/gin"
)

type UpdateAmosCurrencyRequest struct {
	WebServiceAuth *amos.WebServiceAuth `json:"web_service_auth" binding:"required"`
	Date           string               `json:"date" binding:"required"`
}

func NewUpdateAmosCurrencyRequest(c *gin.Context) (*UpdateAmosCurrencyRequest, error) {
	request := &UpdateAmosCurrencyRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		return nil, err
	}

	return request, nil
}

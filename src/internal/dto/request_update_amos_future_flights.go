package dto

import (
	"packages/src/amos"

	"github.com/gin-gonic/gin"
)

type UpdateAmosFutureFlightsRequest struct {
	WebServiceAuth *amos.WebServiceAuth `json:"web_service_auth" binding:"required"`
	Data           []byte               `json:"data" binding:"required"`
}

func NewUpdateAmosFutureFlightsRequest(c *gin.Context) (*UpdateAmosFutureFlightsRequest, error) {
	request := &UpdateAmosFutureFlightsRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		return nil, err
	}

	return request, nil
}

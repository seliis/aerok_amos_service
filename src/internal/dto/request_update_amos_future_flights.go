package dto

import (
	"github.com/gin-gonic/gin"
)

type UpdateAmosFutureFlightsRequest struct {
	Authorization string
	Data          []byte `json:"data" binding:"required"`
}

func NewUpdateAmosFutureFlightsRequest(c *gin.Context) (*UpdateAmosFutureFlightsRequest, error) {
	request := &UpdateAmosFutureFlightsRequest{}

	request.Authorization = c.GetHeader("Authorization")

	if err := c.ShouldBindJSON(request); err != nil {
		return nil, err
	}

	return request, nil
}

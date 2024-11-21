package dto

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UpdateAmosRequest struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Date     string `json:"date" binding:"required"`
}

func NewUpdateAmosRequest(c *gin.Context) (*UpdateAmosRequest, error) {
	request := &UpdateAmosRequest{}

	if err := c.ShouldBindJSON(request); err != nil {
		return nil, err
	}

	fmt.Printf("+%v\n", request)

	return request, nil
}

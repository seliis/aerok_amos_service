package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetID(context *gin.Context) string {
	return context.Param("id")
}

func GetDate(context *gin.Context) (string, error) {
	date := context.Query("date")
	if date == "" {
		return "", errors.New("date is required")
	}

	return date, nil
}

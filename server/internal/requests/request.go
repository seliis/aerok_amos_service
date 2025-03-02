package requests

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func GetID(context *gin.Context) string {
	return context.Param("id")
}

func GetDate(context *gin.Context) (string, error) {
	date, isOk := context.GetQuery("date")
	if !isOk {
		return "", errors.New("date is required")
	}

	return date, nil
}

func GetWorkbook(context *gin.Context, name string) (*excelize.File, error) {
	header, err := context.FormFile(name)
	if err != nil {
		return nil, err
	}

	file, err := header.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return excelize.OpenReader(file)
}

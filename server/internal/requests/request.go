package requests

import "github.com/gin-gonic/gin"

func GetID(context *gin.Context) string {
	return context.Param("id")
}

package handlers

import (
	"encoding/base64"
	"net/http"
	"packages/src/internal/dto"

	"github.com/gin-gonic/gin"
)

type WebServiceHandler struct {
	id       string
	password string
}

func NewWebServiceHandler(id, password string) *WebServiceHandler {
	return &WebServiceHandler{
		id:       id,
		password: password,
	}
}

func (h *WebServiceHandler) GetBasicAuth(c *gin.Context) {
	c.JSON(http.StatusOK, dto.NewSuccessResponse(
		base64.StdEncoding.EncodeToString([]byte(h.id+":"+h.password)),
	))
}

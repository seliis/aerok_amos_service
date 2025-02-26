package server

import (
	"fmt"
	"packages/server/config"
	"packages/server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Start() {
	app := gin.Default()

	setRoutes(app.Group("/api"))

	app.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func setRoutes(api *gin.RouterGroup) {
	{
		h := handlers.NewCurrencyCodeHandler()
		g := api.Group("/currency/code")
		{
			g.POST("/", h.Create)
			g.GET("/:id", h.Read)
			g.PUT("/", h.Update)
			g.DELETE("/:id", h.Delete)
			g.GET("/", h.All)
		}
	}
}

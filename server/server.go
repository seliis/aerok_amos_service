package server

import (
	"fmt"
	"packages/server/config"
	"packages/server/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	setRoutes(app.Group("/api"))

	app.Run(fmt.Sprintf(":%d", config.Server.Port))
}

func setRoutes(api *gin.RouterGroup) {
	setCurrencyRoutes(api.Group("/currency"))
	setExchangeRateRoutes(api.Group("/exchange-rate"))
	setFlightScheduleRoutes(api.Group("/flight-schedule"))
	setAmosRoutes(api.Group("/amos-aim-webservice"))
}

func setCurrencyRoutes(g *gin.RouterGroup) {
	h := handlers.NewCurrencyHandler()

	g.POST("/", h.Create)
	g.GET("/:currency_code", h.Read)
	g.PUT("/", h.Update)
	g.DELETE("/:currency_code", h.Delete)
	g.GET("/", h.All)
	g.POST("/import", h.Import)
}

func setExchangeRateRoutes(g *gin.RouterGroup) {
	h := handlers.NewExchangeRateHandler()

	g.POST("/", h.Create)
	g.GET("/:id", h.Read)
	g.PUT("/", h.Update)
	g.DELETE("/:id", h.Delete)
	g.GET("/", h.All)
	g.PATCH("/", h.UpdateExchangeRates)
	g.GET("/data", h.GetExchangeRates)
}

func setFlightScheduleRoutes(g *gin.RouterGroup) {
	h := handlers.NewFlightScheduleHandler()

	g.POST("/update", h.UpdateFlightSchedules)
}

func setAmosRoutes(g *gin.RouterGroup) {
	h := handlers.NewAmosHandler()

	g.POST("/import-currency", h.ImportCurrency)
	g.POST("/transfer-future-flights", h.TransferFutureFlights)
}

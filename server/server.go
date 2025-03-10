package server

import (
	"fmt"
	"packages/server/config"
	"packages/server/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func Start() error {
	gin.SetMode(gin.ReleaseMode)

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	app.Use(static.Serve("/", static.LocalFile("public", true)))

	setRoutes(app.Group("/api"))

	addr := fmt.Sprintf(":%d", config.Server.Port)

	if err := http3.ListenAndServeTLS(addr, "cert.pem", "key.pem", app); err != nil {
		return err
	}

	return nil
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
	g.GET("/list", h.GetExchangeRates)
	g.GET("/currency", h.GetExchangeRate)
}

func setFlightScheduleRoutes(g *gin.RouterGroup) {
	h := handlers.NewFlightScheduleHandler()

	g.POST("/update", h.UpdateFlightSchedules)
}

func setAmosRoutes(g *gin.RouterGroup) {
	h := handlers.NewAmosHandler()

	g.POST("/import-currency", h.ImportCurrency)
	g.POST("/transfer-future-flights", h.TransferFutureFlights)
	g.POST("/authorize", h.Authorize)
}

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
	protocol := config.Server.Protocol

	switch protocol {
	case 2:
		return startWithHttp2(app, addr, "cert.pem", "key.pem")
	case 3:
		return startWithHttp3(app, addr, "cert.pem", "key.pem")
	default:
		return startWithHttp1(app, addr)
	}
}

func startWithHttp1(app *gin.Engine, addr string) error {
	return app.Run(addr)
}

func startWithHttp2(app *gin.Engine, addr, cert, key string) error {
	return app.RunTLS(addr, cert, key)
}

func startWithHttp3(app *gin.Engine, addr, cert, key string) error {
	return http3.ListenAndServeTLS(addr, cert, key, app)
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

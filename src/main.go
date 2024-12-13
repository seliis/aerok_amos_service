package main

import (
	"fmt"
	"packages/src/config"
	"packages/src/internal/handlers"
	"packages/src/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.Load("config.toml"); err != nil {
		panic(err)
	}

	if err := database.NewClient(); err != nil {
		panic(err)
	}

	if config.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

func getServer() *gin.Engine {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	api := app.Group("/api")
	setRoute(api)

	return app
}

func setRoute(api *gin.RouterGroup) {
	{
		h := handlers.NewCurrencyHandler()

		g := api.Group("/currency")
		g.GET("/currencies", h.GetCurrencies)
		g.GET("/exchangeRate", h.GetExchangeRate)
	}

	{
		basicAuth := config.Amos.BasicAuth

		g := api.Group("/amos", gin.BasicAuth(gin.Accounts{
			basicAuth.Id: basicAuth.Password,
		}))

		h := handlers.NewAmosHandler(basicAuth.Id, basicAuth.Password)
		g.GET("/auth", h.GetBasicAuth)
		g.POST("/updateCurrency/:date", h.UpdateCurrency)
		g.POST("/updateFutureFlights", h.UpdateFutureFlights)
	}
}

func main() {
	server := getServer()

	server.Run(fmt.Sprintf(":%d", config.Server.Port))

	defer func() {
		if err := database.Client.Close(); err != nil {
			panic(err)
		}
	}()
}

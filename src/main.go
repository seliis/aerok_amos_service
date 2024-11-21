package main

import (
	"fmt"
	"packages/src/config"
	"packages/src/internal/handlers"
	"packages/src/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := config.Load("config.toml"); err != nil {
		panic(err)
	}

	if err := database.NewClient(); err != nil {
		panic(err)
	}
}

func main() {
	app := gin.Default()

	if config.Server.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	app.Use(static.Serve("/", static.LocalFile(config.Server.Static, true)))

	setRoute(app.Group("/api"))

	app.Run(fmt.Sprintf(":%d", config.Server.Port))

	defer func() {
		if err := database.Client.Close(); err != nil {
			panic(err)
		}
	}()
}

func setRoute(api *gin.RouterGroup) {
	func() {
		h := handlers.NewCurrencyHandler()

		g := api.Group("/currency")
		g.GET("/exchangeRates", h.GetExchangeRates)
		g.GET("/exchangeRate", h.GetExchangeRate)
		g.GET("/currencies", h.GetCurrencies)
		g.POST("/updateAmos", h.UpdateAmos)
	}()
}

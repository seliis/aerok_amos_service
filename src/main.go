package main

import (
	"fmt"
	"packages/src/config"
	"packages/src/infra/database"
	"packages/src/internal/handlers"

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

	setRoute(app.Group("/api"))

	app.RunTLS(fmt.Sprintf(":%d", config.Server.Port), config.Server.Cert.Pem, config.Server.Cert.Key)

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
		g.GET("/getExchangeRates", h.GetExchangeRates)
		g.POST("/updateAmos", h.UpdateAmos)
	}()
}

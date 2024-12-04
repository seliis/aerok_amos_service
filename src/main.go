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
	setAmos(api)

	return app
}

func setAmos(api *gin.RouterGroup) {
	webService := config.Amos.WebService

	r := api.Group("/amos", gin.BasicAuth(gin.Accounts{
		webService.Id: webService.Password,
	}))

	h := handlers.NewWebServiceHandler(webService.Id, webService.Password)
	r.GET("/auth", h.GetBasicAuth)
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

// func setRoute(api *gin.RouterGroup) {
// 	func() {
// 		h := handlers.NewCurrencyHandler()

// 		g := api.Group("/currency")
// 		g.GET("/exchangeRates", h.GetExchangeRates)
// 		g.GET("/exchangeRate", h.GetExchangeRate)
// 		g.GET("/currencies", h.GetCurrencies)
// 		g.POST("/updateAmos", h.UpdateAmos)
// 	}()

// 	func() {
// 		h := handlers.NewFlightHandler()

// 		g := api.Group("/flight")
// 		g.POST("/updateAmos", h.UpdateAmos)
// 	}()
// }

package main

import (
	"packages/server"
	"packages/server/config"
	"packages/server/database"
	"packages/server/logger"
)

func init() {
	if err := logger.Start("server.log"); err != nil {
		panic(err)
	}

	if err := config.Load("settings.toml"); err != nil {
		panic(err)
	}

	if err := database.Connect(config.Database.URL); err != nil {
		panic(err)
	}

	if err := database.UpsertCurrencies(); err != nil {
		panic(err)
	}
}

func main() {
	if err := server.Start(); err != nil {
		panic(err)
	}

	defer func() {
		if err := database.Disconnect(); err != nil {
			panic(err)
		}
	}()
}

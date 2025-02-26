package main

import (
	"packages/server"
	"packages/server/config"
	"packages/server/database"
)

func init() {
	if err := config.Load("settings.toml"); err != nil {
		panic(err)
	}

	if err := database.Connect(config.Server.Database.URL); err != nil {
		panic(err)
	}
}

func main() {
	server.Start()

	defer func() {
		if err := database.Disconnect(); err != nil {
			panic(err)
		}
	}()
}

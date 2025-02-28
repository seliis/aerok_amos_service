package main

import (
	"io"
	"log"
	"os"
	"packages/server"
	"packages/server/config"
	"packages/server/database"
)

func startLogging(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, f))

	return nil
}

func init() {
	if err := startLogging("server.log"); err != nil {
		panic(err)
	}

	if err := config.Load("settings.toml"); err != nil {
		panic(err)
	}

	if err := database.Connect(config.Database.URL); err != nil {
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

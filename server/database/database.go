package database

import (
	"context"
	"encoding/json"
	"os"
	db "packages/prisma"
	"packages/server/internal/entities"
)

var Client *db.PrismaClient

func Connect(url string) error {
	Client = db.NewClient(db.WithDatasourceURL(url))

	if err := Client.Prisma.Connect(); err != nil {
		return err
	}

	return nil
}

func UpsertCurrencies() error {
	file, err := os.Open("currencies.json")
	if err != nil {
		return err
	}
	defer file.Close()

	var currencies []entities.Currency

	if err := json.NewDecoder(file).Decode(&currencies); err != nil {
		return err
	}

	for _, currency := range currencies {
		if _, err := Client.Currency.UpsertOne(
			db.Currency.Code.Equals(currency.Code),
		).Create(
			db.Currency.Code.Set(currency.Code),
			db.Currency.Name.Set(currency.Name),
			db.Currency.Base.Set(currency.Base),
		).Update(
			db.Currency.Name.Set(currency.Name),
			db.Currency.Base.Set(currency.Base),
		).Exec(context.Background()); err != nil {
			return err
		}
	}

	return nil
}

func Disconnect() error {
	return Client.Prisma.Disconnect()
}

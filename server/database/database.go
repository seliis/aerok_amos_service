package database

import db "packages/prisma"

var Client *db.PrismaClient

func Connect(url string) error {
	Client = db.NewClient(db.WithDatasourceURL(url))

	if err := Client.Prisma.Connect(); err != nil {
		return err
	}

	return nil
}

func Disconnect() error {
	return Client.Prisma.Disconnect()
}

package database_test

import (
	"packages/server/database"
	"testing"
)

func TestConnect(t *testing.T) {
	if err := database.Connect("file:../../prisma/database.db"); err != nil {
		t.Fatal(err)
	}

	t.Logf("Database.Client.Name: %s", database.Client.Name())

	defer func() {
		if err := database.Disconnect(); err != nil {
			t.Fatal(err)
		}
	}()
}

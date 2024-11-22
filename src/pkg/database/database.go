package database

import (
	"database/sql"
	"packages/src/config"

	_ "github.com/mattn/go-sqlite3"
)

var Client *sql.DB

func NewClient() error {
	client, err := sql.Open("sqlite3", "file:database.db?cache=shared")
	if err != nil {
		return err
	}

	if config.Server.Debug {
		if err := dropTables(client); err != nil {
			return err
		}
	}

	if err := createScheme(client); err != nil {
		return err
	}

	Client = client

	return nil
}

func dropTables(client *sql.DB) error {
	if _, err := client.Exec(`
		drop table if exists exchange_rates
	`); err != nil {
		return err
	}

	return nil
}

func createScheme(client *sql.DB) error {
	if _, err := client.Exec(`
        create table if not exists exchange_rates (
			currency_code char(3) not null,
			currency_name varchar(50) not null,
			date char(10) not null,
			direct_rate real not null,
			primary key (currency_code, date)
		)
    `); err != nil {
		return err
	}

	return nil
}

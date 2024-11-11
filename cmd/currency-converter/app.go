package main

import "github.com/rochajg/currency-converter/pkg/database"

func setup() error {
	err := setupDatabase()

	return err
}

func setupDatabase() error {
	return database.InitDB("currencies.db")
}

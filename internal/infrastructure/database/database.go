package database

import "github.com/rochajg/currency-converter/pkg/database"

func InitDB() error {
	return database.InitDB("currencies.db")
}

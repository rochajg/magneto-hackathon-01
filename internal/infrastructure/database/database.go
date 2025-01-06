package database

import (
	"github.com/rochajg/currency-converter/pkg/database"
	"os"
	"path/filepath"
)

func InitDB() error {
	dbPath := getDBPath()

	err := checkDBExists(dbPath)
	if err != nil {
		return err
	}

	return database.InitDB(dbPath)
}

func getDBPath() string {
	pwd := os.Getenv("PWD")

	dbLocation := os.Getenv("DB_LOCATION")
	if dbLocation == "" {
		dbLocation = "/db"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "currencies"
	}

	return pwd + dbLocation + "/" + dbName + ".db"
}

func checkDBExists(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		defer file.Close()
	}

	return nil
}

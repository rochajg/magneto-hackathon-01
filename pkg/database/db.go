package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Conecta ao banco de dados SQLite e inicializa a tabela de taxas de câmbio, se necessário.
func InitDB(dbLocation string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS exchange_rates (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        from_currency TEXT NOT NULL,
        to_currency TEXT NOT NULL,
        rate REAL NOT NULL
    );`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	fmt.Println("Banco de dados SQLite conectado e tabela de taxas de câmbio inicializada.")
	return db, nil
}

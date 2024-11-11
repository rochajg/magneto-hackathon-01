package database

import (
	"database/sql"
	"errors"
	"fmt"
)

type ExchangeRateRepository struct {
	db *sql.DB
}

func NewExchangeRateRepository(db *sql.DB) *ExchangeRateRepository {
	return &ExchangeRateRepository{
		db: db,
	}
}

// Função para adicionar uma nova taxa de câmbio
func (r *ExchangeRateRepository) AddExchangeRate(fromCurrency, toCurrency string, rate float64) error {
	insertQuery := `INSERT INTO exchange_rates (from_currency, to_currency, rate) VALUES (?, ?, ?)`
	_, err := r.db.Exec(insertQuery, fromCurrency, toCurrency, rate)
	if err != nil {
		return err
	}
	fmt.Printf("Taxa de câmbio adicionada: %s -> %s = %.2f\n", fromCurrency, toCurrency, rate)
	return nil
}

// Função para consultar a taxa de câmbio entre duas moedas
func (r *ExchangeRateRepository) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	var rate float64
	query := `SELECT rate FROM exchange_rates WHERE from_currency = ? AND to_currency = ?`
	err := r.db.QueryRow(query, fromCurrency, toCurrency).Scan(&rate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("taxa de câmbio não encontrada")
		}
		return 0, err
	}
	return rate, nil
}

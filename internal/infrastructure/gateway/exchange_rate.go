package gateway

import "github.com/rochajg/currency-converter/pkg/database"

type ExchangeRateDatabaseGateway struct{}

func NewExchangeRateDatabaseGateway() *ExchangeRateDatabaseGateway {
	return &ExchangeRateDatabaseGateway{}
}

func (e *ExchangeRateDatabaseGateway) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	return database.GetExchangeRate(fromCurrency, toCurrency)
}

func (e *ExchangeRateDatabaseGateway) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	return database.AddExchangeRate(fromCurrency, toCurrency, rate)
}

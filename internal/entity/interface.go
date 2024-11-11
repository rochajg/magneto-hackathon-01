package entity

type ExchangeRateRepository interface {
	AddExchangeRate(fromCurrency, toCurrency string, rate float64) error
	GetExchangeRate(fromCurrency, toCurrency string) (float64, error)
}

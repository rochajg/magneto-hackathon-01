package gateway

type ExchangeRateGateway interface {
	GetExchangeRate(fromCurrency, toCurrency string) (float64, error)
	AddExchangeRate(fromCurrency, toCurrency string, rate float64) error
}

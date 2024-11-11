package entity

type ExchangeRate struct {
	FromCurrency string
	ToCurrency   string
	Rate         float64
}

func NewExchangeRate(fromCurrency, toCurrency string, rate float64) *ExchangeRate {
	return &ExchangeRate{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}
}

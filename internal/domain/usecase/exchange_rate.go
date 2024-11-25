package usecase

import "github.com/rochajg/currency-converter/internal/domain/gateway"

type ExchangeRateUseCase struct {
	exchangeRate gateway.ExchangeRateGateway
}

func NewExchangeRateUseCase(exchangeRate gateway.ExchangeRateGateway) *ExchangeRateUseCase {
	return &ExchangeRateUseCase{
		exchangeRate: exchangeRate,
	}
}

func (u *ExchangeRateUseCase) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	return u.exchangeRate.GetExchangeRate(fromCurrency, toCurrency)
}

func (u *ExchangeRateUseCase) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	return u.exchangeRate.AddExchangeRate(fromCurrency, toCurrency, rate)
}

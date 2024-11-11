package usecase

import "github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"

type CreateExchangeRateInput struct {
	FromCurrency string
	ToCurrency   string
	Rate         float64
}

type CreateExchangeRateUseCase struct {
	exchangeRateRepository entity.ExchangeRateRepository
}

func NewCreateExchangeRateUseCase(exchangeRateRepository entity.ExchangeRateRepository) *CreateExchangeRateUseCase {
	return &CreateExchangeRateUseCase{
		exchangeRateRepository: exchangeRateRepository,
	}
}

func (c *CreateExchangeRateUseCase) Execute(input CreateExchangeRateInput) error {
	exchangeRate := entity.NewExchangeRate(input.FromCurrency, input.ToCurrency, input.Rate)
	err := c.exchangeRateRepository.AddExchangeRate(exchangeRate.FromCurrency, exchangeRate.ToCurrency, exchangeRate.Rate)
	if err != nil {
		return err
	}
	return nil
}

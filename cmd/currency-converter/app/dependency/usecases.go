package dependency

import "github.com/rochajg/currency-converter/internal/domain/usecase"

type useCases struct {
	exchangeRateUseCase *usecase.ExchangeRateUseCase
}

func (d *Dependencies) setupUseCases() *Dependencies {
	d.useCases = &useCases{
		exchangeRateUseCase: d.getExchangeRateUseCase(),
	}
	return d
}

func (d *Dependencies) getExchangeRateUseCase() *usecase.ExchangeRateUseCase {
	return usecase.NewExchangeRateUseCase(d.gateways.exchangeRateGateway)
}

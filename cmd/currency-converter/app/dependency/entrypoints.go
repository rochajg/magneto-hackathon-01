package dependency

import (
	"github.com/rochajg/currency-converter/cmd/currency-converter/entrypoint/exchangerate"
	"github.com/rochajg/currency-converter/cmd/currency-converter/entrypoint/healthcheck"
)

type Entrypoints struct {
	HealthCheckEntrypoint  *healthcheck.Entrypoint
	ExchangeRateEntrypoint *exchangerate.Entrypoint
}

func (d *Dependencies) setupEntrypoints() *Dependencies {
	d.Entrypoints = &Entrypoints{
		HealthCheckEntrypoint:  d.getHealthCheckEntrypoint(),
		ExchangeRateEntrypoint: d.getExchangeRateEntrypoint(),
	}
	return d
}

func (d *Dependencies) getHealthCheckEntrypoint() *healthcheck.Entrypoint {
	return healthcheck.NewHealthCheckController()
}

func (d *Dependencies) getExchangeRateEntrypoint() *exchangerate.Entrypoint {
	return exchangerate.NewExchangeRateEntrypoint(*d.useCases.exchangeRateUseCase)
}

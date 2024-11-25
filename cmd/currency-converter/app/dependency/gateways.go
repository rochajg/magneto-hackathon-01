package dependency

import (
	"github.com/rochajg/currency-converter/internal/infrastructure/gateway"
)

type gateways struct {
	exchangeRateGateway *gateway.ExchangeRateDatabaseGateway
}

func (d *Dependencies) setupGateways() *Dependencies {
	d.gateways = &gateways{
		exchangeRateGateway: d.getExchangeRateGateway(),
	}
	return d
}

func (d *Dependencies) getExchangeRateGateway() *gateway.ExchangeRateDatabaseGateway {
	return gateway.NewExchangeRateDatabaseGateway()
}

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/app/dependency"
)

func setupRoutes(router *gin.Engine, e *dependency.Entrypoints) {
	router.GET("/ping", e.HealthCheckEntrypoint.Ping)

	router.POST("/exchange-rate", e.ExchangeRateEntrypoint.AddExchangeRate)
}

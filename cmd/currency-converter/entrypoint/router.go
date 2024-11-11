package entrypoint

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/entrypoint/healthcheck"
)

type Entrypoints struct {
	healthCheckEntrypoint healthcheck.Entrypoint
}

func SetupEntrypoints() *Entrypoints {
	return &Entrypoints{
		healthCheckEntrypoint: *healthcheck.NewHealthCheckController(),
	}
}

func (e *Entrypoints) SetupRoutes(router *gin.Engine) {
	router.GET("/ping", e.healthCheckEntrypoint.Ping)
}

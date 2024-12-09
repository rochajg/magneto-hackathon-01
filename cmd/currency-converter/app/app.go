package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/app/apperror"
	"github.com/rochajg/currency-converter/cmd/currency-converter/app/dependency"
	"github.com/rochajg/currency-converter/cmd/currency-converter/app/middleware"
	"github.com/rochajg/currency-converter/internal/infrastructure/database"
)

func Start(r *gin.Engine) error {
	setupMiddlewares(r)

	deps := dependency.ResolveDependencies()
	setupRoutes(r, deps.Entrypoints)

	err := setupDatabase()
	if err != nil {
		return err
	}

	return nil
}

func setupMiddlewares(r *gin.Engine) {
	r.Use(middleware.RequestID())
	r.Use(apperror.Handler())
}

func setupDatabase() error {
	return database.InitDB()
}

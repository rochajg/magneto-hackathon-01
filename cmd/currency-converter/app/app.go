package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/entrypoint"
	"github.com/rochajg/currency-converter/internal/infrastructure/database"
)

func Start(r *gin.Engine) error {
	setupEntrypoints(r)

	err := setupDatabase()
	if err != nil {
		return err
	}

	return nil
}

func setupEntrypoints(r *gin.Engine) {
	entrypoints := entrypoint.SetupEntrypoints()
	entrypoints.SetupRoutes(r)
}

func setupDatabase() error {
	return database.InitDB()
}

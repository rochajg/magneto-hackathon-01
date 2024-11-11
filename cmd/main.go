package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	repositories "github.com/israelalvesmelo/magneto-hackathon-01/internal/infra/database"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/infra/web/webserver"

	"github.com/israelalvesmelo/magneto-hackathon-01/internal/usecase"
	"github.com/israelalvesmelo/magneto-hackathon-01/pkg/database"
)

func main() {
	// configs, err := LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }

	db, err := database.InitDB("../pkg/database/database.db")
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}

	r := gin.Default()

	exchangeRateRepository := repositories.NewExchangeRateRepository(db)
	createExchangeRateUseCase := usecase.NewCreateExchangeRateUseCase(exchangeRateRepository)
	exchangeHandler := webserver.NewExchangeHandler(*createExchangeRateUseCase)

	r.POST("/exchange-rate", exchangeHandler.AddExchangeRate)
	r.GET("/exchange-rate", exchangeHandler.GetExchangeRate)
	r.POST("/convert", exchangeHandler.ConvertAmount)
	r.Run()
}

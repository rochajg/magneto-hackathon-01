package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/app"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()

	err := app.Start(r)
	if err != nil {
		log.Printf("error initializing app: %v", err)
	}

	if err := r.Run(":" + port); err != nil {
		log.Printf("error: %v", err)
	}
}

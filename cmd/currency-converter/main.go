package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/cmd/currency-converter/router"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := setup()
	if err != nil {
		log.Printf("error initializing app: %v", err)
	}

	r := gin.New()
	router.SetupRoutes(r)

	if err := r.Run(":" + port); err != nil {
		log.Printf("error: %v", err)
	}
}

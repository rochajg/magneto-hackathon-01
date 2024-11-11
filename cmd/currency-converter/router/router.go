package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/ping", ping)
}

func ping(context *gin.Context) {
	log.Println("PONG!")

	context.String(http.StatusOK, "pong")
}

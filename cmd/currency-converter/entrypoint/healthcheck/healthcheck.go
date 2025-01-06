package healthcheck

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Entrypoint struct{}

func NewHealthCheckController() *Entrypoint {
	return &Entrypoint{}
}

func (hc *Entrypoint) Ping(context *gin.Context) {
	log.Println("PONG!")

	context.String(http.StatusOK, "pong")
}

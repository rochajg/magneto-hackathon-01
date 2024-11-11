package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/usecase"
)

type ExchangeHandler struct {
	CreateExchangeRateUseCase usecase.CreateExchangeRateUseCase
}

func NewExchangeHandler(createExchangeRateUseCase usecase.CreateExchangeRateUseCase) *ExchangeHandler {
	return &ExchangeHandler{
		CreateExchangeRateUseCase: createExchangeRateUseCase,
	}
}

func (h *ExchangeHandler) AddExchangeRate(c *gin.Context) {
	var dto usecase.CreateExchangeRateInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.CreateExchangeRateUseCase.Execute(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Taxa de câmbio adicionada com sucesso!"})

}

func (h *ExchangeHandler) GetExchangeRate(c *gin.Context) {
}

func (h *ExchangeHandler) ConvertAmount(c *gin.Context) {
}

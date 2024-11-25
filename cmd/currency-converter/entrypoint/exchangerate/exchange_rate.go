package exchangerate

import (
	"github.com/gin-gonic/gin"
	"github.com/rochajg/currency-converter/internal/domain/usecase"
	"net/http"
	"strconv"
)

type Entrypoint struct {
	useCase usecase.ExchangeRateUseCase
}

func NewExchangeRateEntrypoint(useCase usecase.ExchangeRateUseCase) *Entrypoint {
	return &Entrypoint{
		useCase: useCase,
	}
}

type Request struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Rate         float64 `json:"rate"`
}

func (e *Entrypoint) AddExchangeRate(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := e.useCase.AddExchangeRate(request.FromCurrency, request.ToCurrency, request.Rate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Exchange rate added successfully!",
	})
}

func (e *Entrypoint) GetExchangeRate(c *gin.Context) {
	fromCurrency := c.Query("from")
	toCurrency := c.Query("to")

	rate, err := e.useCase.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rate": rate,
	})
}

func (e *Entrypoint) ConvertCurrency(c *gin.Context) {
	fromCurrency := c.Query("from")
	toCurrency := c.Query("to")
	amount, _ := strconv.ParseFloat(c.Query("amount"), 64)

	convertedAmount, err := e.useCase.ConvertCurrency(fromCurrency, toCurrency, amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"converted_amount": convertedAmount,
	})
}

package apperror

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.Last()
		if err == nil {
			return
		}

		var e *APIError
		switch {
		case errors.As(err.Err, &e):
			c.JSON(e.Status, e)
			return
		default:
			c.JSON(
				http.StatusInternalServerError,
				NewAPIError(
					http.StatusInternalServerError,
					"internal_server_error",
					"Internal Server Error",
					err,
				),
			)
			return
		}
	}
}

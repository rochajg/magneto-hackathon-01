package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// requestIdMiddlewareWriter is a custom response writer that adds the request ID to the response headers
type requestIdMiddlewareWriter struct {
	gin.ResponseWriter
	requestID string
}

// RequestID is a middleware function that adds a request ID to the context
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		c.Writer = &requestIdMiddlewareWriter{
			ResponseWriter: c.Writer,
			requestID:      requestID,
		}

		c.Next()
	}
}

// WriteHeader adds the request ID to the response headers
func (m *requestIdMiddlewareWriter) WriteHeader(statusCode int) {
	m.Header().Add("X-REQUEST-ID", m.requestID)
	m.ResponseWriter.WriteHeader(statusCode)
}

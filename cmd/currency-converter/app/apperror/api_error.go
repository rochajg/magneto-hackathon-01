package apperror

import "fmt"

// APIError implements an apperror interface with some adcional fields
type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Cause   error  `json:"cause"`
}

func NewAPIError(status int, code string, message string, cause error) *APIError {
	return &APIError{
		Message: message,
		Status:  status,
		Code:    code,
		Cause:   cause,
	}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Message: %s;Status: %d;Cause: %v", e.Message, e.Status, e.Cause.Error())
}

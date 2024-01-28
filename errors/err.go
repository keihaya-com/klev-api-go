package errors

import (
	"errors"
	"fmt"
)

type APIError struct {
	Code    string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Details string `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("(%s) %s\n\n%s", e.Code, e.Message, e.Details)
}

func GetError(err error) *APIError {
	var e *APIError
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func IsError(err error, code string) bool {
	apie := GetError(err)
	return apie != nil && apie.Code == code
}

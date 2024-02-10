package klev

import (
	"errors"
	"fmt"
)

type APIError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
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

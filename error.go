package api

import (
	"errors"
	"fmt"
)

type ErrorOut struct {
	ErrorCode string `json:"error,omitempty"`
	Message   string `json:"message,omitempty"`
	Details   string `json:"details,omitempty"`
}

func (e *ErrorOut) Error() string {
	return fmt.Sprintf("(%s) %s\n\n%s", e.ErrorCode, e.Message, e.Details)
}

func GetError(err error) *ErrorOut {
	var e *ErrorOut
	if errors.As(err, &e) {
		return e
	}
	return nil
}

func IsError(err error, code string) bool {
	apie := GetError(err)
	return apie != nil && apie.ErrorCode == code
}

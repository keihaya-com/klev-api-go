package api

import "fmt"

type ErrorOut struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e ErrorOut) Error() string {
	return fmt.Sprintf("(%s) %s", e.Code, e.Message)
}

package api

type ErrorOut struct {
	Error   bool   `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

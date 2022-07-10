package api

import (
	"go.uber.org/zap"
)

type Client struct {
	base baseClient
}

func New(baseURL string, log *zap.SugaredLogger) *Client {
	return &Client{newBase(baseURL, log)}
}

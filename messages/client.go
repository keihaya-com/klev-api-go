package messages

import (
	"github.com/klev-dev/klev-api-go/client"
)

type Client struct {
	H client.HTTP
}

func New(cfg client.Config) *Client {
	return &Client{client.New(cfg)}
}

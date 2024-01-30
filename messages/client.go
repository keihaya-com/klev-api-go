package messages

import "github.com/klev-dev/klev-api-go"

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

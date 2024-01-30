package paths

import (
	"context"

	"github.com/klev-dev/klev-api-go/client"
)

type Client struct {
	H client.HTTP
}

func New(cfg client.Config) *Client {
	return &Client{client.New(cfg)}
}

func (c *Client) Get(ctx context.Context) (map[string]string, error) {
	var out map[string]string
	err := c.H.Get(ctx, "", &out)
	return out, err
}

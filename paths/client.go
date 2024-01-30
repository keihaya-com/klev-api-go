package paths

import (
	"context"

	"github.com/klev-dev/klev-api-go"
)

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

func (c *Client) Get(ctx context.Context) (map[string]string, error) {
	var out map[string]string
	err := c.H.Get(ctx, "", &out)
	return out, err
}

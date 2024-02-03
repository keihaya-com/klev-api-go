package logs

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go"
)

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

func (c *Client) List(ctx context.Context) ([]klev.Log, error) {
	var out klev.Logs
	err := c.H.Get(ctx, fmt.Sprintf("logs"), &out)
	return out.Logs, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.Log, error) {
	var out klev.Logs
	err := c.H.Get(ctx, fmt.Sprintf("logs?metadata=%s", metadata), &out)
	return out.Logs, err
}

func (c *Client) Create(ctx context.Context, in klev.LogCreateParams) (klev.Log, error) {
	var out klev.Log
	err := c.H.Post(ctx, fmt.Sprintf("logs"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.LogID) (klev.Log, error) {
	var out klev.Log
	err := c.H.Get(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

func (c *Client) Stats(ctx context.Context, id klev.LogID) (klev.LogStats, error) {
	var out klev.LogStats
	err := c.H.Get(ctx, fmt.Sprintf("log/%s/stats", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.LogID) (klev.Log, error) {
	var out klev.Log
	err := c.H.Delete(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

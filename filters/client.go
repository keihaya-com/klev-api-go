package filters

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go"
	"github.com/klev-dev/klev-api-go/client"
)

type Client struct {
	H client.HTTP
}

func New(cfg client.Config) *Client {
	return &Client{client.New(cfg)}
}

func (c *Client) List(ctx context.Context) ([]klev.Filter, error) {
	var out klev.Filters
	err := c.H.Get(ctx, fmt.Sprintf("filters"), &out)
	return out.Filters, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.Filter, error) {
	var out klev.Filters
	err := c.H.Get(ctx, fmt.Sprintf("filters?metadata=%s", metadata), &out)
	return out.Filters, err
}

func (c *Client) Create(ctx context.Context, in klev.FilterCreateParams) (klev.Filter, error) {
	var out klev.Filter
	err := c.H.Post(ctx, fmt.Sprintf("filters"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.FilterID) (klev.Filter, error) {
	var out klev.Filter
	err := c.H.Get(ctx, fmt.Sprintf("filter/%s", id), &out)
	return out, err
}

func (c *Client) Status(ctx context.Context, id klev.FilterID) (klev.FilterStatus, error) {
	var out klev.FilterStatus
	err := c.H.Get(ctx, fmt.Sprintf("filter/%s/status", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.FilterID) (klev.Filter, error) {
	var out klev.Filter
	err := c.H.Delete(ctx, fmt.Sprintf("filter/%s", id), &out)
	return out, err
}

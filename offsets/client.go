package offsets

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

func (c *Client) List(ctx context.Context) ([]klev.Offset, error) {
	var out klev.Offsets
	err := c.H.Get(ctx, fmt.Sprintf("offsets"), &out)
	return out.Offsets, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.Offset, error) {
	var out klev.Offsets
	err := c.H.Get(ctx, fmt.Sprintf("offsets?metadata=%s", metadata), &out)
	return out.Offsets, err
}

func (c *Client) Create(ctx context.Context, in klev.OffsetCreateParams) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Post(ctx, fmt.Sprintf("offsets"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.OffsetID) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Get(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

func (c *Client) Set(ctx context.Context, id klev.OffsetID, value int64, valueMetadata string) (klev.Offset, error) {
	return c.SetRaw(ctx, id, klev.OffsetSetParams{
		Value:         value,
		ValueMetadata: valueMetadata,
	})
}

func (c *Client) SetRaw(ctx context.Context, id klev.OffsetID, in klev.OffsetSetParams) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Post(ctx, fmt.Sprintf("offset/%s", id), in, &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.OffsetID) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Delete(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

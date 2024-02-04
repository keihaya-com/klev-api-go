package offsets

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
	return c.UpdateRaw(ctx, id, klev.OffsetUpdateParams{
		Value:         &value,
		ValueMetadata: &valueMetadata,
	})
}

func (c *Client) UpdateRaw(ctx context.Context, id klev.OffsetID, in klev.OffsetUpdateParams) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Patch(ctx, fmt.Sprintf("offset/%s", id), in, &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.OffsetID) (klev.Offset, error) {
	var out klev.Offset
	err := c.H.Delete(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

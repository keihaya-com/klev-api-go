package offsets

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/client"
	"github.com/klev-dev/klev-api-go/logs"
)

type OffsetID string

type Offset struct {
	OffsetID      OffsetID   `json:"offset_id"`
	LogID         logs.LogID `json:"log_id"`
	Metadata      string     `json:"metadata"`
	Value         int64      `json:"value"`
	ValueMetadata string     `json:"value_metadata"`
}

type Offsets struct {
	Offsets []Offset `json:"offsets,omitempty"`
}

type OffsetCreate struct {
	LogID    logs.LogID `json:"log_id"`
	Metadata string     `json:"metadata"`
}

type OffsetSet struct {
	Value         int64  `json:"value"`
	ValueMetadata string `json:"value_metadata"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]Offset, error) {
	var out Offsets
	err := c.H.Get(ctx, fmt.Sprintf("offsets"), &out)
	return out.Offsets, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]Offset, error) {
	var out Offsets
	err := c.H.Get(ctx, fmt.Sprintf("offsets?metadata=%s", metadata), &out)
	return out.Offsets, err
}

func (c *Client) Create(ctx context.Context, in OffsetCreate) (Offset, error) {
	var out Offset
	err := c.H.Post(ctx, fmt.Sprintf("offsets"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.H.Get(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

func (c *Client) Set(ctx context.Context, id OffsetID, value int64, valueMetadata string) (Offset, error) {
	return c.SetRaw(ctx, id, OffsetSet{
		Value:         value,
		ValueMetadata: valueMetadata,
	})
}

func (c *Client) SetRaw(ctx context.Context, id OffsetID, in OffsetSet) (Offset, error) {
	var out Offset
	err := c.H.Post(ctx, fmt.Sprintf("offset/%s", id), in, &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.H.Delete(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

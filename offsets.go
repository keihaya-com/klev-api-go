package api

import (
	"context"
	"fmt"
)

type OffsetID string

type OffsetCreate struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
}

type Offsets struct {
	Offsets []Offset `json:"offsets,omitempty"`
}

type Offset struct {
	OffsetID      OffsetID `json:"offset_id"`
	LogID         LogID    `json:"log_id"`
	Metadata      string   `json:"metadata"`
	Value         int64    `json:"value"`
	ValueMetadata string   `json:"value_metadata"`
}

type OffsetSet struct {
	Value         int64  `json:"value"`
	ValueMetadata string `json:"value_metadata"`
}

func (c *Client) OffsetsList(ctx context.Context) ([]Offset, error) {
	var out Offsets
	err := c.httpGet(ctx, fmt.Sprintf("offsets"), &out)
	return out.Offsets, err
}

func (c *Client) OffsetsFind(ctx context.Context, logID LogID, metadata string) ([]Offset, error) {
	var out Offsets
	err := c.httpGet(ctx, fmt.Sprintf("offsets?log_id=%s&metadata=%s", logID, metadata), &out)
	return out.Offsets, err
}

func (c *Client) OffsetCreate(ctx context.Context, in OffsetCreate) (Offset, error) {
	var out Offset
	err := c.httpPost(ctx, fmt.Sprintf("offsets"), in, &out)
	return out, err
}

func (c *Client) OffsetGet(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.httpGet(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

func (c *Client) OffsetSet(ctx context.Context, id OffsetID, value int64, valueMetadata string) (Offset, error) {
	return c.OffsetSetRaw(ctx, id, OffsetSet{
		Value:         value,
		ValueMetadata: valueMetadata,
	})
}

func (c *Client) OffsetSetRaw(ctx context.Context, id OffsetID, in OffsetSet) (Offset, error) {
	var out Offset
	err := c.httpPost(ctx, fmt.Sprintf("offset/%s", id), in, &out)
	return out, err
}

func (c *Client) OffsetDelete(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.httpDelete(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

package api

import (
	"context"
	"fmt"
)

type OffsetID string

type OffsetCreate struct {
	Metadata string `json:"metadata"`
}

type Offsets struct {
	Offsets []Offset `json:"offsets,omitempty"`
}

type Offset struct {
	OffsetID OffsetID            `json:"offset_id"`
	Metadata string              `json:"metadata"`
	Offsets  map[LogID]LogOffset `json:"offsets"`
}

type LogOffset struct {
	Offset   int64  `json:"offset"`
	Metadata string `json:"metadata"`
}

type OffsetAck struct {
	Offsets map[LogID]LogOffset `json:"offsets"`
}

func (c *Client) OffsetsList(ctx context.Context) ([]Offset, error) {
	var out Offsets
	err := c.httpGet(ctx, fmt.Sprintf("offsets"), &out)
	return out.Offsets, err
}

func (c *Client) OffsetsFind(ctx context.Context, metadata string) ([]Offset, error) {
	var out Offsets
	err := c.httpGet(ctx, fmt.Sprintf("offsets?q=%s", metadata), &out)
	return out.Offsets, err
}

func (c *Client) OffsetCreate(ctx context.Context, in OffsetCreate) (Offset, error) {
	var out Offset
	err := c.httpPost(ctx, fmt.Sprintf("offsets"), in, &out)
	return out, err
}

func (c *Client) OffsetGet(ctx context.Context, id OffsetID, logID LogID) (offset int64, metadata string, err error) {
	off, err := c.OffsetGetAll(ctx, id)
	if err != nil {
		return -1, "", err
	}
	logOff, ok := off.Offsets[logID]
	if !ok {
		return -1, "", &ErrorOut{
			Code:    ErrAPILogOffsetMissing,
			Message: fmt.Sprintf("log '%s' is missing in offset '%s'", logID, id),
		}
	}
	return logOff.Offset, logOff.Metadata, nil
}

func (c *Client) OffsetGetAll(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.httpGet(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

func (c *Client) OffsetAck(ctx context.Context, id OffsetID, logID LogID, offset int64, metadata string) (Offset, error) {
	return c.OffsetAckAll(ctx, id, map[LogID]LogOffset{
		logID: LogOffset{
			Offset:   offset,
			Metadata: metadata,
		}})
}

func (c *Client) OffsetAckAll(ctx context.Context, id OffsetID, in map[LogID]LogOffset) (Offset, error) {
	var out Offset
	err := c.httpPost(ctx, fmt.Sprintf("offset/%s", id), OffsetAck{in}, &out)
	return out, err
}

func (c *Client) OffsetDelete(ctx context.Context, id OffsetID) (Offset, error) {
	var out Offset
	err := c.httpDelete(ctx, fmt.Sprintf("offset/%s", id), &out)
	return out, err
}

package logs

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/client"
)

type LogID string

type Log struct {
	LogID          LogID  `json:"log_id"`
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes,omitempty"`
	TrimSeconds    int64  `json:"trim_seconds,omitempty"`
	CompactSeconds int64  `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64  `json:"expire_seconds,omitempty"`
}

type Logs struct {
	Logs []Log `json:"logs"`
}

type LogCreate struct {
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes"`
	TrimSeconds    int64  `json:"trim_seconds"`
	CompactSeconds int64  `json:"compact_seconds"`
	ExpireSeconds  int64  `json:"expire_seconds"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]Log, error) {
	var out Logs
	err := c.H.Get(ctx, fmt.Sprintf("logs"), &out)
	return out.Logs, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]Log, error) {
	var out Logs
	err := c.H.Get(ctx, fmt.Sprintf("logs?metadata=%s", metadata), &out)
	return out.Logs, err
}

func (c *Client) Create(ctx context.Context, in LogCreate) (Log, error) {
	var out Log
	err := c.H.Post(ctx, fmt.Sprintf("logs"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id LogID) (Log, error) {
	var out Log
	err := c.H.Get(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id LogID) (Log, error) {
	var out Log
	err := c.H.Delete(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

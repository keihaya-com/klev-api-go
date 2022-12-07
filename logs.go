package api

import (
	"context"
	"fmt"
)

type LogID string

type LogIn struct {
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes"`
	TrimSeconds    int64  `json:"trim_seconds"`
	CompactSeconds int64  `json:"compact_seconds"`
	ExpireSeconds  int64  `json:"expire_seconds"`
}

type LogsOut struct {
	Logs []LogOut `json:"logs"`
}

type LogOut struct {
	LogID          LogID  `json:"log_id"`
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes,omitempty"`
	TrimSeconds    int64  `json:"trim_seconds,omitempty"`
	CompactSeconds int64  `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64  `json:"expire_seconds,omitempty"`
}

func (c *Client) LogsList(ctx context.Context) ([]LogOut, error) {
	var out LogsOut
	err := c.HTTPGet(ctx, fmt.Sprintf("logs"), &out)
	return out.Logs, err
}

func (c *Client) LogsFind(ctx context.Context, metadata string) ([]LogOut, error) {
	var out LogsOut
	err := c.HTTPGet(ctx, fmt.Sprintf("logs?q=%s", metadata), &out)
	return out.Logs, err
}

func (c *Client) LogCreate(ctx context.Context, in LogIn) (LogOut, error) {
	var out LogOut
	err := c.HTTPPost(ctx, fmt.Sprintf("logs"), in, &out)
	return out, err
}

func (c *Client) LogGet(ctx context.Context, id LogID) (LogOut, error) {
	var out LogOut
	err := c.HTTPGet(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

func (c *Client) LogDelete(ctx context.Context, id LogID) (LogOut, error) {
	var out LogOut
	err := c.HTTPDelete(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

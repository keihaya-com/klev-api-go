package api

import (
	"context"
	"fmt"
)

type LogID string

type LogCreate struct {
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes"`
	TrimSeconds    int64  `json:"trim_seconds"`
	CompactSeconds int64  `json:"compact_seconds"`
	ExpireSeconds  int64  `json:"expire_seconds"`
}

type Logs struct {
	Logs []Log `json:"logs"`
}

type Log struct {
	LogID          LogID  `json:"log_id"`
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimBytes      int64  `json:"trim_bytes,omitempty"`
	TrimSeconds    int64  `json:"trim_seconds,omitempty"`
	CompactSeconds int64  `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64  `json:"expire_seconds,omitempty"`
}

func (c *Client) LogsList(ctx context.Context) ([]Log, error) {
	var out Logs
	err := c.httpGet(ctx, fmt.Sprintf("logs"), &out)
	return out.Logs, err
}

func (c *Client) LogsFind(ctx context.Context, metadata string) ([]Log, error) {
	var out Logs
	err := c.httpGet(ctx, fmt.Sprintf("logs?metadata=%s", metadata), &out)
	return out.Logs, err
}

func (c *Client) LogCreate(ctx context.Context, in LogCreate) (Log, error) {
	var out Log
	err := c.httpPost(ctx, fmt.Sprintf("logs"), in, &out)
	return out, err
}

func (c *Client) LogGet(ctx context.Context, id LogID) (Log, error) {
	var out Log
	err := c.httpGet(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

func (c *Client) LogDelete(ctx context.Context, id LogID) (Log, error) {
	var out Log
	err := c.httpDelete(ctx, fmt.Sprintf("log/%s", id), &out)
	return out, err
}

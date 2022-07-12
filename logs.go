package api

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

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
	LogID          ksuid.KSUID `json:"log_id"`
	Metadata       string      `json:"metadata"`
	Compacting     bool        `json:"compacting"`
	TrimBytes      int64       `json:"trim_bytes,omitempty"`
	TrimSeconds    int64       `json:"trim_seconds,omitempty"`
	CompactSeconds int64       `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64       `json:"expire_seconds,omitempty"`
}

func (c *Client) LogsList() ([]LogOut, error) {
	var out LogsOut
	err := c.HTTPGet(fmt.Sprintf("logs"), &out)
	return out.Logs, err
}

func (c *Client) LogCreate(in LogIn) (LogOut, error) {
	var out LogOut
	err := c.HTTPPost(fmt.Sprintf("logs"), in, &out)
	return out, err
}

func (c *Client) LogGet(logID ksuid.KSUID) (LogOut, error) {
	var out LogOut
	err := c.HTTPGet(fmt.Sprintf("log/%s", logID), &out)
	return out, err
}

func (c *Client) LogDelete(logID ksuid.KSUID) error {
	var out LogOut
	return c.HTTPDelete(fmt.Sprintf("log/%s", logID), &out)
}

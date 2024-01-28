package filters

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/client"
	"github.com/klev-dev/klev-api-go/logs"
)

type FilterID string

type Filter struct {
	FilterID   FilterID   `json:"filter_id"`
	Metadata   string     `json:"metadata"`
	Source     logs.LogID `json:"source_id"`
	Target     logs.LogID `json:"target_id"`
	Expression string     `json:"expression"`
}

type Filters struct {
	Filters []Filter `json:"filters,omitempty"`
}

type FilterCreate struct {
	Metadata   string     `json:"metadata"`
	SourceID   logs.LogID `json:"source_id"`
	TargetID   logs.LogID `json:"target_id"`
	Expression string     `json:"expression"`
}

type FilterStatus struct {
	FilterID FilterID `json:"filter_id"`

	Active         bool   `json:"active"`
	InactiveReason string `json:"inactive_reason,omitempty"`

	AvailableOffset int64 `json:"available_offset"`

	DeliverOffset int64  `json:"deliver_offset"`
	DeliverTime   int64  `json:"deliver_time,omitempty"`
	DeliverError  string `json:"deliver_error,omitempty"`

	NextDeliverOffset int64 `json:"next_deliver_offset"`
	NextDeliverTime   int64 `json:"next_deliver_time,omitempty"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]Filter, error) {
	var out Filters
	err := c.H.Get(ctx, fmt.Sprintf("filters"), &out)
	return out.Filters, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]Filter, error) {
	var out Filters
	err := c.H.Get(ctx, fmt.Sprintf("filters?metadata=%s", metadata), &out)
	return out.Filters, err
}

func (c *Client) Create(ctx context.Context, in FilterCreate) (Filter, error) {
	var out Filter
	err := c.H.Post(ctx, fmt.Sprintf("filters"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id FilterID) (Filter, error) {
	var out Filter
	err := c.H.Get(ctx, fmt.Sprintf("filter/%s", id), &out)
	return out, err
}

func (c *Client) Status(ctx context.Context, id FilterID) (FilterStatus, error) {
	var out FilterStatus
	err := c.H.Get(ctx, fmt.Sprintf("filter/%s/status", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id FilterID) (Filter, error) {
	var out Filter
	err := c.H.Delete(ctx, fmt.Sprintf("filter/%s", id), &out)
	return out, err
}

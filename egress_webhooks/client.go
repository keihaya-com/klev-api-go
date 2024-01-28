package egress_webhooks

import (
	"context"
	"fmt"
	"time"

	"github.com/klev-dev/klev-api-go/client"
	"github.com/klev-dev/klev-api-go/logs"
)

type EgressWebhookID string

type EgressWebhook struct {
	WebhookID   EgressWebhookID `json:"webhook_id"`
	Metadata    string          `json:"metadata"`
	LogID       logs.LogID      `json:"log_id"`
	Destination string          `json:"destination"`
	Payload     string          `json:"payload"`
	Secret      string          `json:"secret,omitempty"`
}

type EgressWebhooks struct {
	Items []EgressWebhook `json:"items"`
}

type EgressWebhookCreate struct {
	Metadata    string     `json:"metadata"`
	LogID       logs.LogID `json:"log_id"`
	Destination string     `json:"destination"`
	Payload     string     `json:"payload"`
}

type EgressWebhookRotate struct {
	ExpireSeconds int64 `json:"expire_seconds"`
}

type EgressWebhookStatus struct {
	WebhookID EgressWebhookID `json:"webhook_id"`

	Active         bool   `json:"active"`
	InactiveReason string `json:"inactive_reason,omitempty"`

	AvailableOffset int64 `json:"available_offset"`

	DeliverOffset int64  `json:"deliver_offset"`
	DeliverTime   int64  `json:"deliver_time,omitempty"`
	DeliverResp   string `json:"deliver_resp,omitempty"`
	DeliverError  string `json:"deliver_error,omitempty"`

	NextDeliverOffset int64 `json:"next_deliver_offset"`
	NextDeliverTime   int64 `json:"next_deliver_time,omitempty"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]EgressWebhook, error) {
	var out EgressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]EgressWebhook, error) {
	var out EgressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) Create(ctx context.Context, in EgressWebhookCreate) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.H.Post(ctx, fmt.Sprintf("egress_webhooks"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id EgressWebhookID) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) Rotate(ctx context.Context, id EgressWebhookID, expireDuration time.Duration) (EgressWebhook, error) {
	var in = EgressWebhookRotate{int64(expireDuration.Seconds())}
	return c.RotateRaw(ctx, id, in)
}

func (c *Client) RotateRaw(ctx context.Context, id EgressWebhookID, in EgressWebhookRotate) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.H.Patch(ctx, fmt.Sprintf("egress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) Status(ctx context.Context, id EgressWebhookID) (EgressWebhookStatus, error) {
	var out EgressWebhookStatus
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhook/%s/status", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id EgressWebhookID) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.H.Delete(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

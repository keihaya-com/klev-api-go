package api

import (
	"context"
	"fmt"
	"time"
)

type EgressWebhookID string

type EgressWebhookCreate struct {
	Metadata    string `json:"metadata"`
	LogID       LogID  `json:"log_id"`
	Destination string `json:"destination"`
}

type EgressWebhooks struct {
	Items []EgressWebhook `json:"items"`
}

type EgressWebhook struct {
	WebhookID   EgressWebhookID `json:"webhook_id"`
	Metadata    string          `json:"metadata"`
	LogID       LogID           `json:"log_id"`
	Destination string          `json:"destination"`
	Secret      string          `json:"secret,omitempty"`
}

type EgressWebhookRotate struct {
	ExpireSeconds int64 `json:"expire_seconds"`
}

type EgressWebhookStatus struct {
	WebhookID EgressWebhookID `json:"webhook_id"`

	Active         bool   `json:"active"`
	InactiveReason string `json:"inactive_reason,omitempty"`

	AvailableOffset int64 `json:"available_offset"`
	NextOffset      int64 `json:"next_offset"`

	DeliverOffset   int64  `json:"deliver_offset"`
	DeliverTime     int64  `json:"deliver_time,omitempty"`
	DeliverResp     string `json:"deliver_resp,omitempty"`
	DeliverError    string `json:"deliver_error,omitempty"`
	NextDeliverTime int64  `json:"next_deliver_time,omitempty"`
}

func (c *Client) EgressWebhooksList(ctx context.Context) ([]EgressWebhook, error) {
	var out EgressWebhooks
	err := c.httpGet(ctx, fmt.Sprintf("egress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) EgressWebhooksFind(ctx context.Context, metadata string) ([]EgressWebhook, error) {
	var out EgressWebhooks
	err := c.httpGet(ctx, fmt.Sprintf("egress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) EgressWebhookCreate(ctx context.Context, in EgressWebhookCreate) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.httpPost(ctx, fmt.Sprintf("egress_webhooks"), in, &out)
	return out, err
}

func (c *Client) EgressWebhookGet(ctx context.Context, id EgressWebhookID) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.httpGet(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) EgressWebhookDelete(ctx context.Context, id EgressWebhookID) (EgressWebhook, error) {
	var out EgressWebhook
	err := c.httpDelete(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) EgressWebhookRotate(ctx context.Context, id EgressWebhookID, expireDuration time.Duration) (EgressWebhook, error) {
	var in = EgressWebhookRotate{int64(expireDuration.Seconds())}
	var out EgressWebhook
	err := c.httpPatch(ctx, fmt.Sprintf("egress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) EgressWebhookStatus(ctx context.Context, id EgressWebhookID) (EgressWebhookStatus, error) {
	var out EgressWebhookStatus
	err := c.httpGet(ctx, fmt.Sprintf("egress_webhook/%s/status", id), &out)
	return out, err
}

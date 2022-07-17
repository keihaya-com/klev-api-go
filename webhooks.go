package api

import (
	"context"
	"fmt"

	"github.com/segmentio/ksuid"
)

type WebhookIn struct {
	LogID    ksuid.KSUID `json:"log_id"`
	Metadata string      `json:"metadata"`
	Type     string      `json:"type"`
	Secret   string      `json:"secret"`
}

type WebhooksOut struct {
	Webhooks []WebhookOut `json:"webhooks"`
}

type WebhookOut struct {
	WebhookID ksuid.KSUID `json:"webhook_id"`
	LogID     ksuid.KSUID `json:"log_id"`
	Metadata  string      `json:"metadata"`
	Type      string      `json:"type"`
}

func (c *Client) WebhooksList(ctx context.Context) ([]WebhookOut, error) {
	var out WebhooksOut
	err := c.HTTPGet(ctx, fmt.Sprintf("webhooks"), &out)
	return out.Webhooks, err
}

func (c *Client) WebhookCreate(ctx context.Context, in WebhookIn) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPPost(ctx, fmt.Sprintf("webhooks"), in, &out)
	return out, err
}

func (c *Client) WebhookGet(ctx context.Context, webhookID ksuid.KSUID) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPGet(ctx, fmt.Sprintf("webhook/%s", webhookID), &out)
	return out, err
}

func (c *Client) WebhookDelete(ctx context.Context, webhookID ksuid.KSUID) error {
	var out WebhookOut
	return c.HTTPDelete(ctx, fmt.Sprintf("webhook/%s", webhookID), &out)
}

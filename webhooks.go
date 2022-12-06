package api

import (
	"context"
	"fmt"
)

type WebhookID string

type WebhookIn struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
	Type     string `json:"type"`
	Secret   string `json:"secret"`
}

type WebhooksOut struct {
	Webhooks []WebhookOut `json:"webhooks"`
}

type WebhookOut struct {
	WebhookID WebhookID `json:"webhook_id"`
	LogID     LogID     `json:"log_id"`
	Metadata  string    `json:"metadata"`
	Type      string    `json:"type"`
}

func (c *Client) WebhooksList(ctx context.Context) ([]WebhookOut, error) {
	var out WebhooksOut
	err := c.HTTPGet(ctx, fmt.Sprintf("webhooks"), &out)
	return out.Webhooks, err
}

func (c *Client) WebhooksFind(ctx context.Context, metadata string) ([]WebhookOut, error) {
	var out WebhooksOut
	err := c.HTTPGet(ctx, fmt.Sprintf("webhooks?q=%s", metadata), &out)
	return out.Webhooks, err
}

func (c *Client) WebhookCreate(ctx context.Context, in WebhookIn) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPPost(ctx, fmt.Sprintf("webhooks"), in, &out)
	return out, err
}

func (c *Client) WebhookGet(ctx context.Context, id WebhookID) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPGet(ctx, fmt.Sprintf("webhook/%s", id), &out)
	return out, err
}

func (c *Client) WebhookDelete(ctx context.Context, id WebhookID) error {
	var out WebhookOut
	return c.HTTPDelete(ctx, fmt.Sprintf("webhook/%s", id), &out)
}

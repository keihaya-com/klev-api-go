package api

import (
	"context"
	"fmt"
)

type WebhookID string

type WebhookCreate struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
	Type     string `json:"type"`
	Secret   string `json:"secret"`
}

type Webhooks struct {
	Webhooks []Webhook `json:"webhooks"`
}

type Webhook struct {
	WebhookID WebhookID `json:"webhook_id"`
	LogID     LogID     `json:"log_id"`
	Metadata  string    `json:"metadata"`
	Type      string    `json:"type"`
}

func (c *Client) WebhooksList(ctx context.Context) ([]Webhook, error) {
	var out Webhooks
	err := c.httpGet(ctx, fmt.Sprintf("webhooks"), &out)
	return out.Webhooks, err
}

func (c *Client) WebhooksFind(ctx context.Context, metadata string) ([]Webhook, error) {
	var out Webhooks
	err := c.httpGet(ctx, fmt.Sprintf("webhooks?metadata=%s", metadata), &out)
	return out.Webhooks, err
}

func (c *Client) WebhookCreate(ctx context.Context, in WebhookCreate) (Webhook, error) {
	var out Webhook
	err := c.httpPost(ctx, fmt.Sprintf("webhooks"), in, &out)
	return out, err
}

func (c *Client) WebhookGet(ctx context.Context, id WebhookID) (Webhook, error) {
	var out Webhook
	err := c.httpGet(ctx, fmt.Sprintf("webhook/%s", id), &out)
	return out, err
}

func (c *Client) WebhookDelete(ctx context.Context, id WebhookID) (Webhook, error) {
	var out Webhook
	err := c.httpDelete(ctx, fmt.Sprintf("webhook/%s", id), &out)
	return out, err
}

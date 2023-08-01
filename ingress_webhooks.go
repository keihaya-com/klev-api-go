package api

import (
	"context"
	"fmt"
)

type IngressWebhookID string

type IngressWebhookCreate struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
	Type     string `json:"type"`
	Secret   string `json:"secret"`
}

type IngressWebhooks struct {
	Items []IngressWebhook `json:"items"`
}

type IngressWebhook struct {
	WebhookID IngressWebhookID `json:"webhook_id"`
	Metadata  string           `json:"metadata"`
	LogID     LogID            `json:"log_id"`
	Type      string           `json:"type"`
}

type IngressWebhookRotate struct {
	Secret string `json:"secret"`
}

func (c *Client) IngressWebhooksList(ctx context.Context) ([]IngressWebhook, error) {
	var out IngressWebhooks
	err := c.httpGet(ctx, fmt.Sprintf("ingress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) IngressWebhooksFind(ctx context.Context, metadata string) ([]IngressWebhook, error) {
	var out IngressWebhooks
	err := c.httpGet(ctx, fmt.Sprintf("ingress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) IngressWebhookCreate(ctx context.Context, in IngressWebhookCreate) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.httpPost(ctx, fmt.Sprintf("ingress_webhooks"), in, &out)
	return out, err
}

func (c *Client) IngressWebhookGet(ctx context.Context, id IngressWebhookID) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.httpGet(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) IngressWebhookRotate(ctx context.Context, id IngressWebhookID, secret string) (IngressWebhook, error) {
	return c.IngressWebhookRotateRaw(ctx, id, IngressWebhookRotate{secret})
}

func (c *Client) IngressWebhookRotateRaw(ctx context.Context, id IngressWebhookID, in IngressWebhookRotate) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.httpPatch(ctx, fmt.Sprintf("ingress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) IngressWebhookDelete(ctx context.Context, id IngressWebhookID) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.httpDelete(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

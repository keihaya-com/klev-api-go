package ingress_webhooks

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/client"
	"github.com/klev-dev/klev-api-go/logs"
)

type IngressWebhookID string

type IngressWebhook struct {
	WebhookID IngressWebhookID `json:"webhook_id"`
	Metadata  string           `json:"metadata"`
	LogID     logs.LogID       `json:"log_id"`
	Type      string           `json:"type"`
}

type IngressWebhooks struct {
	Items []IngressWebhook `json:"items"`
}

type CreateParams struct {
	LogID    logs.LogID `json:"log_id"`
	Metadata string     `json:"metadata"`
	Type     string     `json:"type"`
	Secret   string     `json:"secret"`
}

type RotateParams struct {
	Secret string `json:"secret"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]IngressWebhook, error) {
	var out IngressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]IngressWebhook, error) {
	var out IngressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) Create(ctx context.Context, in CreateParams) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.H.Post(ctx, fmt.Sprintf("ingress_webhooks"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id IngressWebhookID) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) Rotate(ctx context.Context, id IngressWebhookID, secret string) (IngressWebhook, error) {
	return c.RotateRaw(ctx, id, RotateParams{secret})
}

func (c *Client) RotateRaw(ctx context.Context, id IngressWebhookID, in RotateParams) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.H.Patch(ctx, fmt.Sprintf("ingress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id IngressWebhookID) (IngressWebhook, error) {
	var out IngressWebhook
	err := c.H.Delete(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

package ingress_webhooks

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go"
)

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

func (c *Client) List(ctx context.Context) ([]klev.IngressWebhook, error) {
	var out klev.IngressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.IngressWebhook, error) {
	var out klev.IngressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) Create(ctx context.Context, in klev.IngressWebhookCreateParams) (klev.IngressWebhook, error) {
	var out klev.IngressWebhook
	err := c.H.Post(ctx, fmt.Sprintf("ingress_webhooks"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.IngressWebhookID) (klev.IngressWebhook, error) {
	var out klev.IngressWebhook
	err := c.H.Get(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) Rotate(ctx context.Context, id klev.IngressWebhookID, secret string) (klev.IngressWebhook, error) {
	return c.RotateRaw(ctx, id, klev.IngressWebhookRotateParams{secret})
}

func (c *Client) RotateRaw(ctx context.Context, id klev.IngressWebhookID, in klev.IngressWebhookRotateParams) (klev.IngressWebhook, error) {
	var out klev.IngressWebhook
	err := c.H.Patch(ctx, fmt.Sprintf("ingress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.IngressWebhookID) (klev.IngressWebhook, error) {
	var out klev.IngressWebhook
	err := c.H.Delete(ctx, fmt.Sprintf("ingress_webhook/%s", id), &out)
	return out, err
}

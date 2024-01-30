package egress_webhooks

import (
	"context"
	"fmt"
	"time"

	"github.com/klev-dev/klev-api-go"
)

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

func (c *Client) List(ctx context.Context) ([]klev.EgressWebhook, error) {
	var out klev.EgressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhooks"), &out)
	return out.Items, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.EgressWebhook, error) {
	var out klev.EgressWebhooks
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhooks?metadata=%s", metadata), &out)
	return out.Items, err
}

func (c *Client) Create(ctx context.Context, in klev.EgressWebhookCreateParams) (klev.EgressWebhook, error) {
	var out klev.EgressWebhook
	err := c.H.Post(ctx, fmt.Sprintf("egress_webhooks"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.EgressWebhookID) (klev.EgressWebhook, error) {
	var out klev.EgressWebhook
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

func (c *Client) Rotate(ctx context.Context, id klev.EgressWebhookID, expireDuration time.Duration) (klev.EgressWebhook, error) {
	var in = klev.EgressWebhookRotateParams{int64(expireDuration.Seconds())}
	return c.RotateRaw(ctx, id, in)
}

func (c *Client) RotateRaw(ctx context.Context, id klev.EgressWebhookID, in klev.EgressWebhookRotateParams) (klev.EgressWebhook, error) {
	var out klev.EgressWebhook
	err := c.H.Patch(ctx, fmt.Sprintf("egress_webhook/%s/secret", id), in, &out)
	return out, err
}

func (c *Client) Status(ctx context.Context, id klev.EgressWebhookID) (klev.EgressWebhookStatus, error) {
	var out klev.EgressWebhookStatus
	err := c.H.Get(ctx, fmt.Sprintf("egress_webhook/%s/status", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.EgressWebhookID) (klev.EgressWebhook, error) {
	var out klev.EgressWebhook
	err := c.H.Delete(ctx, fmt.Sprintf("egress_webhook/%s", id), &out)
	return out, err
}

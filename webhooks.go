package api

import (
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

func (c *Client) WebhooksList() ([]WebhookOut, error) {
	var out WebhooksOut
	err := c.HTTPGet(fmt.Sprintf("webhooks"), &out)
	return out.Webhooks, err
}

func (c *Client) WebhookCreate(in WebhookIn) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPPost(fmt.Sprintf("webhooks"), in, &out)
	return out, err
}

func (c *Client) WebhookGet(webhookID ksuid.KSUID) (WebhookOut, error) {
	var out WebhookOut
	err := c.HTTPGet(fmt.Sprintf("webhook/%s", webhookID), &out)
	return out, err
}

func (c *Client) WebhookDelete(webhookID ksuid.KSUID) error {
	var out WebhookOut
	return c.HTTPDelete(fmt.Sprintf("webhook/%s", webhookID), &out)
}

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

func (k *Client) WebhooksList() ([]WebhookOut, error) {
	var out WebhooksOut
	err := k.base.Get(fmt.Sprintf("webhooks"), &out)
	return out.Webhooks, err
}

func (k *Client) WebhookCreate(in WebhookIn) (WebhookOut, error) {
	var out WebhookOut
	err := k.base.Post(fmt.Sprintf("webhooks"), in, &out)
	return out, err
}

func (k *Client) WebhookGet(webhookID ksuid.KSUID) error {
	var out WebhookOut
	return k.base.Get(fmt.Sprintf("webhook/%s", webhookID), &out)
}

func (k *Client) WebhookDelete(webhookID ksuid.KSUID) error {
	var out WebhookOut
	return k.base.Delete(fmt.Sprintf("webhook/%s", webhookID), &out)
}

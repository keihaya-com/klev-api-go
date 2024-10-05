package klev

type IngressWebhook struct {
	WebhookID IngressWebhookID   `json:"webhook_id"`
	Metadata  string             `json:"metadata,omitempty"`
	LogID     LogID              `json:"log_id"`
	Type      IngressWebhookType `json:"type"`
}

type IngressWebhooks struct {
	IngressWebhooks []IngressWebhook `json:"ingress_webhooks"`
}

type IngressWebhookCreateParams struct {
	LogID    LogID              `json:"log_id"`
	Metadata string             `json:"metadata,omitempty"`
	Type     IngressWebhookType `json:"type"`
	Secret   string             `json:"secret"`
}

type IngressWebhookUpdateParams struct {
	Metadata *string `json:"metadata,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

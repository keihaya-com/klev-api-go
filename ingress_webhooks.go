package klev

type IngressWebhookID string

type IngressWebhook struct {
	WebhookID IngressWebhookID `json:"webhook_id"`
	Metadata  string           `json:"metadata"`
	LogID     LogID            `json:"log_id"`
	Type      string           `json:"type"`
}

type IngressWebhooks struct {
	Items []IngressWebhook `json:"items"`
}

type IngressWebhookCreateParams struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
	Type     string `json:"type"`
	Secret   string `json:"secret"`
}

type IngressWebhookRotateParams struct {
	Secret string `json:"secret"`
}

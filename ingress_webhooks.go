package klev

type IngressWebhookID string

func ParseIngressWebhookID(id string) (IngressWebhookID, error) {
	if err := validate(id, "iwh"); err != nil {
		return IngressWebhookID(""), err
	}
	return IngressWebhookID(id), nil
}

type IngressWebhook struct {
	WebhookID IngressWebhookID `json:"webhook_id"`
	Metadata  string           `json:"metadata,omitempty"`
	LogID     LogID            `json:"log_id"`
	Type      string           `json:"type"`
}

type IngressWebhooks struct {
	Items []IngressWebhook `json:"items"`
}

type IngressWebhookCreateParams struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata,omitempty"`
	Type     string `json:"type"`
	Secret   string `json:"secret"`
}

type IngressWebhookUpdateParams struct {
	Metadata *string `json:"metadata,omitempty"`
	Secret   *string `json:"secret,omitempty"`
}

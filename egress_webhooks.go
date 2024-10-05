package klev

type EgressWebhook struct {
	WebhookID   EgressWebhookID      `json:"webhook_id"`
	Metadata    string               `json:"metadata"`
	LogID       LogID                `json:"log_id"`
	Destination string               `json:"destination"`
	Payload     EgressWebhookPayload `json:"payload"`
	Secret      string               `json:"secret,omitempty"`
}

type EgressWebhooks struct {
	EgressWebhooks []EgressWebhook `json:"egress_webhooks"`
}

type EgressWebhookCreateParams struct {
	Metadata    string               `json:"metadata,omitempty"`
	LogID       LogID                `json:"log_id"`
	Destination string               `json:"destination"`
	Payload     EgressWebhookPayload `json:"payload"`
}

type EgressWebhookUpdateParams struct {
	Metadata    *string `json:"metadata,omitempty"`
	Destination *string `json:"destination,omitempty"`
}

type EgressWebhookRotateParams struct {
	ExpireSeconds int64 `json:"expire_seconds"`
}

type EgressWebhookStatus struct {
	WebhookID EgressWebhookID `json:"webhook_id"`

	Active         bool   `json:"active"`
	InactiveReason string `json:"inactive_reason,omitempty"`

	AvailableOffset int64 `json:"available_offset"`

	DeliverOffset int64  `json:"deliver_offset"`
	DeliverTime   int64  `json:"deliver_time,omitempty"`
	DeliverResp   string `json:"deliver_resp,omitempty"`
	DeliverError  string `json:"deliver_error,omitempty"`

	NextDeliverOffset int64 `json:"next_deliver_offset"`
	NextDeliverTime   int64 `json:"next_deliver_time,omitempty"`
}

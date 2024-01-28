// Code generated by 'make gen-errors'; DO NOT EDIT
package egress_webhooks

import "github.com/klev-dev/klev-api-go/errors"

const (
	ErrEgressWebhookPathInvalid       = "ERR_KLEV_EGRESS_WEBHOOKS_API_0001"
	ErrEgressWebhookLogIDFieldInvalid = "ERR_KLEV_EGRESS_WEBHOOKS_API_0002"
)

func IsErrEgressWebhookPathInvalid(err error) bool {
	return errors.IsError(err, ErrEgressWebhookPathInvalid)
}

func IsErrEgressWebhookLogIDFieldInvalid(err error) bool {
	return errors.IsError(err, ErrEgressWebhookLogIDFieldInvalid)
}

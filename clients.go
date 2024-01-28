package api

import (
	"github.com/klev-dev/klev-api-go/client"
	"github.com/klev-dev/klev-api-go/egress_webhooks"
	"github.com/klev-dev/klev-api-go/filters"
	"github.com/klev-dev/klev-api-go/ingress_webhooks"
	"github.com/klev-dev/klev-api-go/logs"
	"github.com/klev-dev/klev-api-go/messages"
	"github.com/klev-dev/klev-api-go/offsets"
	"github.com/klev-dev/klev-api-go/paths"
	"github.com/klev-dev/klev-api-go/tokens"
)

// Clients wraps interactions with klev api
type Clients struct {
	EgressWebhooks  *egress_webhooks.Client
	Filters         *filters.Client
	IngressWebhooks *ingress_webhooks.Client
	Logs            *logs.Client
	Messages        *messages.Client
	Offsets         *offsets.Client
	Paths           *paths.Client
	Tokens          *tokens.Client
}

// New create a new clients from a config
func New(cfg client.Config) *Clients {
	c := client.New(cfg)
	return &Clients{
		EgressWebhooks:  &egress_webhooks.Client{c},
		Filters:         &filters.Client{c},
		IngressWebhooks: &ingress_webhooks.Client{c},
		Logs:            &logs.Client{c},
		Messages:        &messages.Client{c},
		Offsets:         &offsets.Client{c},
		Paths:           &paths.Client{c},
		Tokens:          &tokens.Client{c},
	}
}

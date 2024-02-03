package klev

import (
	"fmt"
)

type Subject string

const (
	SubjectEgressWebhooks  Subject = "egress_webhooks"
	SubjectFilters         Subject = "filters"
	SubjectIngressWebhooks Subject = "ingress_webhooks"
	SubjectLogs            Subject = "logs"
	SubjectMessages        Subject = "messages"
	SubjectOffsets         Subject = "offsets"
	SubjectTokens          Subject = "tokens"
)

var AllSubjects = []Subject{
	SubjectEgressWebhooks, SubjectFilters,
	SubjectIngressWebhooks,
	SubjectLogs, SubjectMessages,
	SubjectOffsets, SubjectTokens,
}

type Action string

const (
	// Generic
	ActionList   Action = "list"
	ActionCreate Action = "create"
	ActionGet    Action = "get"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"

	// Messages
	ActionPublish Action = "publish"
	ActionConsume Action = "consume"
	ActionCleanup Action = "cleanup"

	// Webhooks
	ActionRotate Action = "rotate"
	ActionStatus Action = "status"
)

var defaultActions = []Action{
	ActionList,
	ActionCreate,
	ActionGet,
	ActionDelete,
}

func (s Subject) Actions() []Action {
	switch s {
	case SubjectMessages:
		return []Action{ActionPublish, ActionConsume, ActionCleanup}
	case SubjectOffsets:
		return append(defaultActions, ActionUpdate)
	case SubjectIngressWebhooks:
		return append(defaultActions, ActionRotate)
	case SubjectEgressWebhooks:
		return append(defaultActions, ActionRotate, ActionStatus)
	case SubjectFilters:
		return append(defaultActions, ActionStatus)
	}
	return defaultActions
}

func ACLSubject(subject Subject) string {
	return fmt.Sprintf("%s", subject)
}

func ACLAction(subject Subject, action Action) string {
	return fmt.Sprintf("%s:%s", subject, action)
}

func ACLObject(subject Subject, action Action, id string) string {
	return fmt.Sprintf("%s:%s:%s", subject, action, id)
}

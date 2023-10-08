package api

import (
	"fmt"
)

type Subject string

var (
	SubjectLogs            Subject = "logs"
	SubjectMessages        Subject = "messages"
	SubjectOffsets         Subject = "offsets"
	SubjectTokens          Subject = "tokens"
	SubjectIngressWebhooks Subject = "ingress_webhooks"
	SubjectEgressWebhooks  Subject = "egress_webhooks"
)

var AllSubjects = []Subject{
	SubjectLogs, SubjectMessages,
	SubjectOffsets, SubjectTokens,
	SubjectIngressWebhooks, SubjectEgressWebhooks,
}

type Action string

var (
	// Generic
	ActionList   Action = "list"
	ActionCreate Action = "create"
	ActionGet    Action = "get"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"

	// Messages
	ActionPublish Action = "publish"
	ActionConsume Action = "consume"

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
		return []Action{ActionPublish, ActionConsume}
	case SubjectOffsets:
		return append(defaultActions, ActionUpdate)
	case SubjectIngressWebhooks:
		return append(defaultActions, ActionRotate)
	case SubjectEgressWebhooks:
		return append(defaultActions, ActionRotate, ActionStatus)
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

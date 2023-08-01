package api

import (
	"fmt"
)

type Subject string

var (
	SubjectLogs            Subject = "logs"
	SubjectOffsets         Subject = "offsets"
	SubjectTokens          Subject = "tokens"
	SubjectIngressWebhooks Subject = "ingress_webhooks"
	SubjectEgressWebhooks  Subject = "egress_webhooks"
)

var AllSubjects = []Subject{
	SubjectLogs, SubjectOffsets, SubjectTokens,
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

	// Egress webhooks
	ActionRotate Action = "rotate"
	ActionStatus Action = "status"
)

func (s Subject) Actions() []Action {
	var actions = []Action{
		ActionList,
		ActionCreate,
		ActionGet,
		ActionUpdate,
		ActionDelete,
	}
	switch s {
	case SubjectLogs:
		actions = append(actions, ActionPublish, ActionConsume)
	case SubjectEgressWebhooks:
		actions = append(actions, ActionRotate, ActionStatus)
	case SubjectIngressWebhooks:
		actions = append(actions, ActionRotate)
	}
	return actions
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

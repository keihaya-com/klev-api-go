package api

import "fmt"

type Subject string

var (
	SubjectLogs     Subject = "logs"
	SubjectOffsets  Subject = "offsets"
	SubjectTokens   Subject = "tokens"
	SubjectWebhooks Subject = "webhooks"
)

type Action string

var (
	ActionList   Action = "list"
	ActionCreate Action = "create"
	ActionGet    Action = "get"
	ActionUpdate Action = "update"
	ActionDelete Action = "delete"

	ActionPublish Action = "publish"
	ActionConsume Action = "consume"
)

func ACLSubject(subject Subject) string {
	return fmt.Sprintf("%s", subject)
}

func ACLAction(subject Subject, action Action) string {
	return fmt.Sprintf("%s:%s", subject, action)
}

func ACLObject(subject Subject, action Action, id string) string {
	return fmt.Sprintf("%s:%s:%s", subject, action, id)
}

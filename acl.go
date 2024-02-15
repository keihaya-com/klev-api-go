package klev

import (
	"strings"
)

type ACLItem struct {
	Subject
	Action
	Object string
}

func (l ACLItem) MarshalText() ([]byte, error) {
	var parts = []string{l.Subject.string}
	if l.Action.string != "" {
		parts = append(parts, l.Action.string)
		if l.Object != "" {
			parts = append(parts, l.Object)
		}
	}
	return []byte(strings.Join(parts, ":")), nil
}

func (l *ACLItem) UnmarshalText(text []byte) error {
	item := string(text)

	subjectStr, item, found := strings.Cut(item, ":")
	subject, err := ParseSubject(subjectStr)
	if err != nil {
		return err
	}
	l.Subject = subject
	if !found {
		return nil
	}

	actionStr, item, found := strings.Cut(item, ":")
	action, err := ParseAction(actionStr)
	if err != nil {
		return err
	}
	l.Action = action
	if !found {
		return nil
	}

	// TODO validate object
	if item == "" {
		return ErrACLObjectMissing()
	}
	l.Object = item
	return nil
}

func ACLSubject(subject Subject) ACLItem {
	return ACLItem{Subject: subject}
}

func ACLAction(subject Subject, action Action) ACLItem {
	return ACLItem{Subject: subject, Action: action}
}

func ACLObject(subject Subject, action Action, id interface{ IDValue() string }) ACLItem {
	return ACLItem{subject, action, id.IDValue()}
}

package klev

type Filter struct {
	FilterID   FilterID `json:"filter_id"`
	Metadata   string   `json:"metadata"`
	Source     LogID    `json:"source_id"`
	Target     LogID    `json:"target_id"`
	Expression string   `json:"expression"`
}

type Filters struct {
	Filters []Filter `json:"filters"`
}

type FilterCreateParams struct {
	Metadata   string `json:"metadata,omitempty"`
	SourceID   LogID  `json:"source_id"`
	TargetID   LogID  `json:"target_id"`
	Expression string `json:"expression"`
}

type FilterUpdateParams struct {
	Metadata   *string `json:"metadata,omitempty"`
	Expression *string `json:"expression,omitempty"`
}

type FilterStatus struct {
	FilterID FilterID `json:"filter_id"`

	Active         bool   `json:"active"`
	InactiveReason string `json:"inactive_reason,omitempty"`

	AvailableOffset int64 `json:"available_offset"`

	DeliverOffset int64  `json:"deliver_offset"`
	DeliverTime   int64  `json:"deliver_time,omitempty"`
	DeliverError  string `json:"deliver_error,omitempty"`

	NextDeliverOffset int64 `json:"next_deliver_offset"`
	NextDeliverTime   int64 `json:"next_deliver_time,omitempty"`
}

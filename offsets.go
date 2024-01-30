package klev

type OffsetID string

type Offset struct {
	OffsetID      OffsetID `json:"offset_id"`
	LogID         LogID    `json:"log_id"`
	Metadata      string   `json:"metadata"`
	Value         int64    `json:"value"`
	ValueMetadata string   `json:"value_metadata"`
}

type Offsets struct {
	Offsets []Offset `json:"offsets,omitempty"`
}

type OffsetCreateParams struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata"`
}

type OffsetSetParams struct {
	Value         int64  `json:"value"`
	ValueMetadata string `json:"value_metadata"`
}

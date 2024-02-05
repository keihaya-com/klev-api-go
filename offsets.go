package klev

type OffsetID string

func ParseOffsetID(id string) (OffsetID, error) {
	if err := validate(id, "off"); err != nil {
		return OffsetID(""), err
	}
	return OffsetID(id), nil
}

type Offset struct {
	OffsetID      OffsetID `json:"offset_id"`
	LogID         LogID    `json:"log_id"`
	Metadata      string   `json:"metadata"`
	Value         int64    `json:"value"`
	ValueMetadata string   `json:"value_metadata"`
}

type Offsets struct {
	Offsets []Offset `json:"offsets"`
}

type OffsetCreateParams struct {
	LogID    LogID  `json:"log_id"`
	Metadata string `json:"metadata,omitempty"`
}

type OffsetUpdateParams struct {
	Metadata      *string `json:"metadata,omitempty"`
	Value         *int64  `json:"value,omitempty"`
	ValueMetadata *string `json:"value_metadata,omitempty"`
}

package klev

type LogID string

type Log struct {
	LogID          LogID  `json:"log_id"`
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimSeconds    int64  `json:"trim_seconds,omitempty"`
	TrimSize       int64  `json:"trim_size,omitempty"`
	TrimCount      int64  `json:"trim_count,omitempty"`
	CompactSeconds int64  `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64  `json:"expire_seconds,omitempty"`
}

type Logs struct {
	Logs []Log `json:"logs"`
}

type LogCreateParams struct {
	Metadata       string `json:"metadata"`
	Compacting     bool   `json:"compacting"`
	TrimSeconds    int64  `json:"trim_seconds"`
	TrimSize       int64  `json:"trim_size"`
	TrimCount      int64  `json:"trim_count"`
	CompactSeconds int64  `json:"compact_seconds"`
	ExpireSeconds  int64  `json:"expire_seconds"`
}

type LogStats struct {
	Size  int64 `json:"size"`
	Count int64 `json:"count"`
}

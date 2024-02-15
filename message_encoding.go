package klev

import (
	"encoding/base64"
	"time"
)

func (e MessageEncoding) EncodeTime(t time.Time) int64 {
	return t.UTC().UnixMicro()
}

func (e MessageEncoding) EncodeTimeOpt(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}

	ts := e.EncodeTime(t)
	return &ts
}

func (e MessageEncoding) DecodeTime(ts int64) time.Time {
	return time.UnixMicro(ts).UTC()
}

func (e MessageEncoding) DecodeTimeOpt(ts *int64) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return e.DecodeTime(*ts)
}

func (e MessageEncoding) EncodeData(val []byte) *string {
	if val == nil {
		return nil
	}
	var s string
	switch e {
	case MessageEncodingString:
		s = string(val)
	case MessageEncodingBase64:
		s = base64.StdEncoding.EncodeToString(val)
	}
	return &s
}

func (e MessageEncoding) DecodeData(val *string) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	switch e {
	case MessageEncodingString:
		return []byte(*val), nil
	case MessageEncodingBase64:
		return base64.StdEncoding.DecodeString(*val)
	}
	return nil, ErrMessageEncodingInvalid(e.string, "string, base64")
}

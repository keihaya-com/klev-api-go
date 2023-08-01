package api

import (
	"encoding/base64"
	"fmt"
	"time"
)

type MessageEncoding interface {
	String() string

	EncodeData(b []byte) *string
	DecodeData(s *string) ([]byte, error)

	EncodeTime(t time.Time) int64
	EncodeTimeOpt(t time.Time) *int64
	DecodeTime(ts int64) time.Time
	DecodeTimeOpt(ts *int64) time.Time
}

type encodingCommon struct{}

func (e encodingCommon) EncodeTime(t time.Time) int64 {
	return t.UTC().UnixMicro()
}

func (e encodingCommon) EncodeTimeOpt(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}

	ts := e.EncodeTime(t)
	return &ts
}

func (e encodingCommon) DecodeTime(ts int64) time.Time {
	return time.UnixMicro(ts).UTC()
}

func (e encodingCommon) DecodeTimeOpt(ts *int64) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return e.DecodeTime(*ts)
}

type encodingString struct {
	encodingCommon
}

func (c encodingString) String() string {
	return "string"
}

func (c encodingString) EncodeData(val []byte) *string {
	if val == nil {
		return nil
	}
	s := string(val)
	return &s
}

func (c encodingString) DecodeData(val *string) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	return []byte(*val), nil
}

var EncodingString encodingString
var _ MessageEncoding = EncodingString

type encodingBase64 struct {
	encodingCommon
}

func (c encodingBase64) String() string {
	return "base64"
}

func (c encodingBase64) EncodeData(val []byte) *string {
	if val == nil {
		return nil
	}
	s := base64.StdEncoding.EncodeToString(val)
	return &s
}

func (c encodingBase64) DecodeData(val *string) ([]byte, error) {
	if val == nil {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(*val)
}

var EncodingBase64 encodingBase64
var _ MessageEncoding = EncodingBase64

func parseEncoding(s string) (MessageEncoding, error) {
	switch s {
	case "base64":
		return EncodingBase64, nil
	case "string":
		return EncodingString, nil
	default:
		return EncodingBase64, fmt.Errorf("unknown encoding: %s", s)
	}
}

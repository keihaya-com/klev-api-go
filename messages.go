package klev

import (
	"time"
)

const (
	// OffsetOldest represents the smallest offset still available
	// Use it to consume all messages, starting at the beginning of the log
	OffsetOldest int64 = -2
	// OffsetNewest represents the offset that will be used for the next produce
	// Use it to consume messages, starting from the next one produced
	OffsetNewest int64 = -1
	// OffsetInvalid is the offset returned when error is detected
	OffsetInvalid int64 = -3
)

type PublishIn struct {
	Encoding MessageEncoding    `json:"encoding,omitempty"`
	Messages []PublishMessageIn `json:"messages,omitempty"`
}

type PublishMessageIn struct {
	Time  *int64  `json:"time,omitempty"`
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

type PublishOut struct {
	NextOffset int64 `json:"next_offset"`
}

type PublishMessage struct {
	Time  time.Time
	Key   []byte
	Value []byte
}

func NewPublishMessage(key, value string) PublishMessage {
	return PublishMessage{Key: []byte(key), Value: []byte(value)}
}

func NewPublishMessageKey(key string) PublishMessage {
	return PublishMessage{Key: []byte(key)}
}

func NewPublishMessageValue(value string) PublishMessage {
	return PublishMessage{Value: []byte(value)}
}

type PostIn struct {
	Encoding MessageEncoding `json:"encoding,omitempty"`
	Time     *int64          `json:"time,omitempty"`
	Key      *string         `json:"key,omitempty"`
	Value    *string         `json:"value,omitempty"`
}

type PostOut struct {
	NextOffset int64 `json:"next_offset"`
}

type ConsumeOpts struct {
	Offset   *int64
	OffsetID *OffsetID
	Size     *int32
	Poll     *time.Duration
	Encoding MessageEncoding
}

type ConsumeOpt func(opts ConsumeOpts) ConsumeOpts

func ConsumeOffset(offset int64) ConsumeOpt {
	return func(opts ConsumeOpts) ConsumeOpts {
		opts.Offset = &offset
		return opts
	}
}

func ConsumeOldest() ConsumeOpt {
	return ConsumeOffset(OffsetOldest)
}

func ConsumeNewest() ConsumeOpt {
	return ConsumeOffset(OffsetNewest)
}

func ConsumeOffsetID(offsetID OffsetID) ConsumeOpt {
	return func(opts ConsumeOpts) ConsumeOpts {
		opts.OffsetID = &offsetID
		return opts
	}
}

func ConsumeLen(sz int32) ConsumeOpt {
	return func(opts ConsumeOpts) ConsumeOpts {
		opts.Size = &sz
		return opts
	}
}

func ConsumePoll(d time.Duration) ConsumeOpt {
	return func(opts ConsumeOpts) ConsumeOpts {
		opts.Poll = &d
		return opts
	}
}

func ConsumeEncoding(enc MessageEncoding) ConsumeOpt {
	return func(opts ConsumeOpts) ConsumeOpts {
		opts.Encoding = enc
		return opts
	}
}

type ConsumeOut struct {
	NextOffset int64               `json:"next_offset"`
	Encoding   MessageEncoding     `json:"encoding,omitempty"`
	Messages   []ConsumeMessageOut `json:"messages,omitempty"`
}

type ConsumeMessageOut struct {
	Offset int64   `json:"offset"`
	Time   int64   `json:"time"`
	Key    *string `json:"key,omitempty"`
	Value  *string `json:"value,omitempty"`
}

type ConsumeMessage struct {
	Offset int64
	Time   time.Time
	Key    []byte
	Value  []byte
}

func (out ConsumeMessageOut) Decode(coder MessageEncoding) (ConsumeMessage, error) {
	k, err := coder.DecodeData(out.Key)
	if err != nil {
		return ConsumeMessage{}, err
	}
	v, err := coder.DecodeData(out.Value)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return ConsumeMessage{
		Offset: out.Offset,
		Time:   coder.DecodeTime(out.Time),
		Key:    k,
		Value:  v,
	}, nil
}

type GetOut struct {
	Encoding MessageEncoding `json:"encoding,omitempty"`
	Offset   int64           `json:"offset"`
	Time     int64           `json:"time"`
	Key      *string         `json:"key,omitempty"`
	Value    *string         `json:"value,omitempty"`
}

func (out GetOut) Decode() (ConsumeMessage, error) {
	k, err := out.Encoding.DecodeData(out.Key)
	if err != nil {
		return ConsumeMessage{}, err
	}
	v, err := out.Encoding.DecodeData(out.Value)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return ConsumeMessage{
		Offset: out.Offset,
		Time:   out.Encoding.DecodeTime(out.Time),
		Key:    k,
		Value:  v,
	}, nil
}

type GetByKeyIn struct {
	Encoding MessageEncoding `json:"encoding,omitempty"`
	Key      *string         `json:"key,omitempty"`
}

type CleanupOpt func(opts CleanupIn) CleanupIn

func CleanupTrimAge(age time.Duration) CleanupOpt {
	return func(opts CleanupIn) CleanupIn {
		opts.TrimSeconds = int64(age / time.Second)
		return opts
	}
}

func CleanupTrimSize(size int64) CleanupOpt {
	return func(opts CleanupIn) CleanupIn {
		opts.TrimSize = size
		return opts
	}
}

func CleanupTrimCount(count int64) CleanupOpt {
	return func(opts CleanupIn) CleanupIn {
		opts.TrimCount = count
		return opts
	}
}

func CleanupCompactAge(age time.Duration) CleanupOpt {
	return func(opts CleanupIn) CleanupIn {
		opts.CompactSeconds = int64(age / time.Second)
		return opts
	}
}

func CleanupExpireAge(age time.Duration) CleanupOpt {
	return func(opts CleanupIn) CleanupIn {
		opts.ExpireSeconds = int64(age / time.Second)
		return opts
	}
}

type CleanupIn struct {
	TrimSeconds    int64 `json:"trim_seconds,omitempty"`
	TrimSize       int64 `json:"trim_size,omitempty"`
	TrimCount      int64 `json:"trim_count,omitempty"`
	CompactSeconds int64 `json:"compact_seconds,omitempty"`
	ExpireSeconds  int64 `json:"expire_seconds,omitempty"`
}

type CleanupOut struct {
	Size int64 `json:"size"`
}

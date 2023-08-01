package api

import (
	"context"
	"fmt"
	"strings"
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

type ConsumeOut struct {
	NextOffset int64               `json:"next_offset"`
	Encoding   string              `json:"encoding"`
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

type GetOut struct {
	Encoding string  `json:"encoding"`
	Offset   int64   `json:"offset"`
	Time     int64   `json:"time"`
	Key      *string `json:"key,omitempty"`
	Value    *string `json:"value,omitempty"`
}

type GetByKeyIn struct {
	Encoding string  `json:"encoding"`
	Key      *string `json:"key"`
}

type consumeOpts struct {
	offset   *int64
	offsetID *OffsetID
	sz       *int32
	poll     *time.Duration
	encoding MessageEncoding
}

func (c consumeOpts) query() string {
	var params []string
	if c.offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *c.offset))
	}
	if c.offsetID != nil {
		params = append(params, fmt.Sprintf("offset_id=%s", *c.offsetID))
	}
	if c.sz != nil {
		params = append(params, fmt.Sprintf("len=%d", *c.sz))
	}
	if c.poll != nil {
		params = append(params, fmt.Sprintf("poll=%d", (*c.poll)/time.Millisecond))
	}
	params = append(params, fmt.Sprintf("encoding=%s", c.encoding))
	return strings.Join(params, "&")
}

type ConsumeOpt func(opts consumeOpts) consumeOpts

func ConsumeOffset(offset int64) ConsumeOpt {
	return func(opts consumeOpts) consumeOpts {
		opts.offset = &offset
		return opts
	}
}

func ConsumeOffsetID(offsetID OffsetID) ConsumeOpt {
	return func(opts consumeOpts) consumeOpts {
		opts.offsetID = &offsetID
		return opts
	}
}

func ConsumeLen(sz int32) ConsumeOpt {
	return func(opts consumeOpts) consumeOpts {
		opts.sz = &sz
		return opts
	}
}

func ConsumePoll(d time.Duration) ConsumeOpt {
	return func(opts consumeOpts) consumeOpts {
		opts.poll = &d
		return opts
	}
}

func ConsumeEncoding(enc MessageEncoding) ConsumeOpt {
	return func(opts consumeOpts) consumeOpts {
		opts.encoding = enc
		return opts
	}
}

func (c *Client) Consume(ctx context.Context, id LogID, opts ...ConsumeOpt) (int64, []ConsumeMessage, error) {
	copts := consumeOpts{encoding: EncodingBase64}
	for _, opt := range opts {
		copts = opt(copts)
	}

	var out ConsumeOut
	err := c.httpGet(ctx, fmt.Sprintf("messages/%s?%s", id, copts.query()), &out)
	if err != nil {
		return 0, nil, err
	}

	coder, err := ParseMessageEncoding(out.Encoding)
	if err != nil {
		return 0, nil, err
	}

	var msgs []ConsumeMessage
	for _, msg := range out.Messages {
		k, err := coder.DecodeData(msg.Key)
		if err != nil {
			return 0, nil, err
		}
		v, err := coder.DecodeData(msg.Value)
		if err != nil {
			return 0, nil, err
		}

		msgs = append(msgs, ConsumeMessage{
			Offset: msg.Offset,
			Time:   coder.DecodeTime(msg.Time),
			Key:    k,
			Value:  v,
		})
	}

	return out.NextOffset, msgs, err
}

func (c *Client) Get(ctx context.Context, id LogID, offset int64) (ConsumeMessage, error) {
	var out GetOut
	err := c.httpGet(ctx, fmt.Sprintf("message/%s?offset=%d&encoding=base64", id, offset), &out)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return out.Decode()
}

func (c *Client) GetByKey(ctx context.Context, id LogID, key []byte) (ConsumeMessage, error) {
	coder := EncodingBase64
	var out GetOut
	err := c.httpPost(ctx, fmt.Sprintf("message/%s/key", id), GetByKeyIn{
		Encoding: coder.String(),
		Key:      coder.EncodeData(key),
	}, &out)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return out.Decode()
}

func (out GetOut) Decode() (ConsumeMessage, error) {
	coder, err := ParseMessageEncoding(out.Encoding)
	if err != nil {
		return ConsumeMessage{}, err
	}

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

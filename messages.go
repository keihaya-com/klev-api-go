package api

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type PublishIn struct {
	Encoding string             `json:"encoding"`
	Messages []PublishMessageIn `json:"messages"`
}

type PublishMessageIn struct {
	Time  *int64  `json:"time"`
	Key   *string `json:"key"`
	Value *string `json:"value"`
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

func (c *Client) Publish(ctx context.Context, id LogID, messages []PublishMessage) (int64, error) {
	coder := EncodingBase64
	in := PublishIn{
		Encoding: coder.String(),
	}
	for _, msg := range messages {
		in.Messages = append(in.Messages, PublishMessageIn{
			Time:  coder.EncodeTimeOpt(msg.Time),
			Key:   coder.EncodeData(msg.Key),
			Value: coder.EncodeData(msg.Value),
		})
	}

	return c.PublishRaw(ctx, id, in)
}

func (c *Client) PublishRaw(ctx context.Context, id LogID, in PublishIn) (int64, error) {
	var out PublishOut
	err := c.httpPost(ctx, fmt.Sprintf("messages/%s", id), in, &out)
	return out.NextOffset, err
}

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

func ConsumeOldest() ConsumeOpt {
	return ConsumeOffset(OffsetOldest)
}

func ConsumeNewest() ConsumeOpt {
	return ConsumeOffset(OffsetNewest)
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

	var msgs = make([]ConsumeMessage, len(out.Messages))
	for i, outMsg := range out.Messages {
		msg, err := outMsg.Decode(coder)
		if err != nil {
			return 0, nil, err
		}

		msgs[i] = msg
	}

	return out.NextOffset, msgs, err
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

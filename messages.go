package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/segmentio/ksuid"
)

type ConsumeOut struct {
	NextOffset int64               `json:"next_offset"`
	Encoding   string              `json:"encoding"`
	Messages   []ConsumeMessageOut `json:"messages,omitempty"`
}

type ConsumeMessageOut struct {
	Offset int64     `json:"offset"`
	Time   time.Time `json:"time"`
	Key    *string   `json:"key,omitempty"`
	Value  *string   `json:"value,omitempty"`
}

type PublishIn struct {
	Encoding string             `json:"encoding"`
	Messages []PublishMessageIn `json:"messages"`
}

type PublishMessageIn struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}

type PublishOut struct {
	NextOffset int64 `json:"next_offset"`
}

type GetOut struct {
	Encoding string    `json:"encoding"`
	Offset   int64     `json:"offset"`
	Time     time.Time `json:"time"`
	Key      *string   `json:"key,omitempty"`
	Value    *string   `json:"value,omitempty"`
}

type PostIn struct {
	Encoding string  `json:"encoding"`
	Key      *string `json:"key"`
	Value    *string `json:"value"`
}

type PostOut struct {
	NextOffset int64 `json:"next_offset"`
}

type PublishMessage struct {
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

func (c *Client) Publish(ctx context.Context, logID ksuid.KSUID, messages []PublishMessage) (int64, error) {
	in := PublishIn{
		Encoding: "base64",
	}
	for _, msg := range messages {
		in.Messages = append(in.Messages, PublishMessageIn{
			Key:   encodeBase64(msg.Key),
			Value: encodeBase64(msg.Value),
		})
	}

	return c.PublishRaw(ctx, logID, in)
}

func (c *Client) PublishRaw(ctx context.Context, logID ksuid.KSUID, in PublishIn) (int64, error) {
	var out PublishOut
	err := c.HTTPPost(ctx, fmt.Sprintf("messages/%s", logID), in, &out)
	return out.NextOffset, err
}

func (c *Client) Post(ctx context.Context, logID ksuid.KSUID, key []byte, value []byte) (int64, error) {
	in := PostIn{
		Encoding: "base64",
		Key:      encodeBase64(key),
		Value:    encodeBase64(value),
	}

	return c.PostRaw(ctx, logID, in)
}

func (c *Client) PostRaw(ctx context.Context, logID ksuid.KSUID, in PostIn) (int64, error) {
	var out PostOut
	err := c.HTTPPost(ctx, fmt.Sprintf("message/%s", logID), in, &out)
	return out.NextOffset, err
}

func encodeBase64(b []byte) *string {
	if b == nil {
		return nil
	}
	s := base64.StdEncoding.EncodeToString(b)
	return &s
}

func encodeLiteral(b []byte) *string {
	if b == nil {
		return nil
	}
	s := string(b)
	return &s
}

type ConsumeMessage struct {
	Offset int64
	Time   time.Time
	Key    []byte
	Value  []byte
}

func (c *Client) Consume(ctx context.Context, logID ksuid.KSUID, offset int64, sz int32) (int64, []ConsumeMessage, error) {
	var out ConsumeOut
	err := c.HTTPGet(ctx, fmt.Sprintf("messages/%s?offset=%d&len=%d&encoding=base64", logID, offset, sz), &out)
	if err != nil {
		return 0, nil, err
	}

	var decoder = decodeBase64
	if out.Encoding == "string" {
		decoder = decodeLiteral
	}

	var msgs []ConsumeMessage
	for _, msg := range out.Messages {
		k, err := decoder(msg.Key)
		if err != nil {
			return 0, nil, err
		}
		v, err := decoder(msg.Value)
		if err != nil {
			return 0, nil, err
		}

		msgs = append(msgs, ConsumeMessage{
			Offset: msg.Offset,
			Time:   msg.Time,
			Key:    k,
			Value:  v,
		})
	}

	return out.NextOffset, msgs, err
}

func (c *Client) Get(ctx context.Context, logID ksuid.KSUID, offset int64) (ConsumeMessage, error) {
	var out GetOut
	err := c.HTTPGet(ctx, fmt.Sprintf("message/%s?offset=%d&encoding=base64", logID, offset), &out)
	if err != nil {
		return ConsumeMessage{}, err
	}

	var decoder = decodeBase64
	if out.Encoding == "string" {
		decoder = decodeLiteral
	}

	k, err := decoder(out.Key)
	if err != nil {
		return ConsumeMessage{}, err
	}
	v, err := decoder(out.Value)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return ConsumeMessage{
		Offset: out.Offset,
		Time:   out.Time,
		Key:    k,
		Value:  v,
	}, err
}

func decodeBase64(s *string) ([]byte, error) {
	if s == nil {
		return nil, nil
	}
	return base64.StdEncoding.DecodeString(*s)
}

func decodeLiteral(s *string) ([]byte, error) {
	if s == nil {
		return nil, nil
	}
	return []byte(*s), nil
}

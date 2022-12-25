package api

import (
	"context"
	"encoding/base64"
	"fmt"
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

type PostIn struct {
	Encoding string  `json:"encoding"`
	Time     *int64  `json:"time"`
	Key      *string `json:"key"`
	Value    *string `json:"value"`
}

type PostOut struct {
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
	in := PublishIn{
		Encoding: "base64",
	}
	for _, msg := range messages {
		in.Messages = append(in.Messages, PublishMessageIn{
			Time:  encodeTime(msg.Time),
			Key:   encodeBase64(msg.Key),
			Value: encodeBase64(msg.Value),
		})
	}

	return c.PublishRaw(ctx, id, in)
}

func (c *Client) PublishRaw(ctx context.Context, id LogID, in PublishIn) (int64, error) {
	var out PublishOut
	err := c.httpPost(ctx, fmt.Sprintf("messages/%s", id), in, &out)
	return out.NextOffset, err
}

func (c *Client) Post(ctx context.Context, id LogID, t time.Time, key []byte, value []byte) (int64, error) {
	in := PostIn{
		Encoding: "base64",
		Time:     encodeTime(t),
		Key:      encodeBase64(key),
		Value:    encodeBase64(value),
	}

	return c.PostRaw(ctx, id, in)
}

func (c *Client) PostRaw(ctx context.Context, id LogID, in PostIn) (int64, error) {
	var out PostOut
	err := c.httpPost(ctx, fmt.Sprintf("message/%s", id), in, &out)
	return out.NextOffset, err
}

func encodeTime(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	ts := t.UnixMicro()
	return &ts
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

package api

import (
	"context"
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

func (c *Client) Post(ctx context.Context, id LogID, t time.Time, key []byte, value []byte) (int64, error) {
	coder := EncodingBase64
	in := PostIn{
		Encoding: coder.String(),
		Time:     coder.EncodeTimeOpt(t),
		Key:      coder.EncodeData(key),
		Value:    coder.EncodeData(value),
	}

	return c.PostRaw(ctx, id, in)
}

func (c *Client) PostRaw(ctx context.Context, id LogID, in PostIn) (int64, error) {
	var out PostOut
	err := c.httpPost(ctx, fmt.Sprintf("message/%s", id), in, &out)
	return out.NextOffset, err
}

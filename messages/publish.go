package messages

import (
	"context"
	"fmt"
	"time"

	"github.com/klev-dev/klev-api-go"
)

func (c *Client) Publish(ctx context.Context, id klev.LogID, messages []klev.PublishMessage) (int64, error) {
	coder := klev.EncodingBase64
	in := klev.PublishIn{
		Encoding: coder.String(),
	}
	for _, msg := range messages {
		in.Messages = append(in.Messages, klev.PublishMessageIn{
			Time:  coder.EncodeTimeOpt(msg.Time),
			Key:   coder.EncodeData(msg.Key),
			Value: coder.EncodeData(msg.Value),
		})
	}

	return c.PublishRaw(ctx, id, in)
}

func (c *Client) PublishRaw(ctx context.Context, id klev.LogID, in klev.PublishIn) (int64, error) {
	var out klev.PublishOut
	err := c.H.Post(ctx, fmt.Sprintf("messages/%s", id), in, &out)
	return out.NextOffset, err
}

func (c *Client) Post(ctx context.Context, id klev.LogID, t time.Time, key []byte, value []byte) (int64, error) {
	coder := klev.EncodingBase64
	in := klev.PostIn{
		Encoding: coder.String(),
		Time:     coder.EncodeTimeOpt(t),
		Key:      coder.EncodeData(key),
		Value:    coder.EncodeData(value),
	}

	return c.PostRaw(ctx, id, in)
}

func (c *Client) PostRaw(ctx context.Context, id klev.LogID, in klev.PostIn) (int64, error) {
	var out klev.PostOut
	err := c.H.Post(ctx, fmt.Sprintf("message/%s", id), in, &out)
	return out.NextOffset, err
}

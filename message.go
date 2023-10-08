package api

import (
	"context"
	"fmt"
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

type PostIn struct {
	Encoding string  `json:"encoding"`
	Time     *int64  `json:"time"`
	Key      *string `json:"key"`
	Value    *string `json:"value"`
}

type PostOut struct {
	NextOffset int64 `json:"next_offset"`
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

type GetOut struct {
	Encoding string  `json:"encoding"`
	Offset   int64   `json:"offset"`
	Time     int64   `json:"time"`
	Key      *string `json:"key,omitempty"`
	Value    *string `json:"value,omitempty"`
}

func (c *Client) GetByOffset(ctx context.Context, id LogID, offset int64) (ConsumeMessage, error) {
	var out GetOut
	err := c.httpGet(ctx, fmt.Sprintf("message/%s?offset=%d&encoding=base64", id, offset), &out)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return out.Decode()
}

type GetByKeyIn struct {
	Encoding string  `json:"encoding"`
	Key      *string `json:"key"`
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

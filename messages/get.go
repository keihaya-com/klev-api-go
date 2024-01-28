package messages

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/logs"
)

type GetOut struct {
	Encoding string  `json:"encoding"`
	Offset   int64   `json:"offset"`
	Time     int64   `json:"time"`
	Key      *string `json:"key,omitempty"`
	Value    *string `json:"value,omitempty"`
}

func (c *Client) GetByOffset(ctx context.Context, id logs.LogID, offset int64) (ConsumeMessage, error) {
	var out GetOut
	err := c.H.Get(ctx, fmt.Sprintf("message/%s?offset=%d&encoding=base64", id, offset), &out)
	if err != nil {
		return ConsumeMessage{}, err
	}

	return out.Decode()
}

type GetByKeyIn struct {
	Encoding string  `json:"encoding"`
	Key      *string `json:"key"`
}

func (c *Client) GetByKey(ctx context.Context, id logs.LogID, key []byte) (ConsumeMessage, error) {
	coder := EncodingBase64
	var out GetOut
	err := c.H.Post(ctx, fmt.Sprintf("message/%s/key", id), GetByKeyIn{
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

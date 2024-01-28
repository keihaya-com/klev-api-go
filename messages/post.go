package messages

import (
	"context"
	"fmt"
	"time"

	"github.com/klev-dev/klev-api-go/logs"
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

func (c *Client) Post(ctx context.Context, id logs.LogID, t time.Time, key []byte, value []byte) (int64, error) {
	coder := EncodingBase64
	in := PostIn{
		Encoding: coder.String(),
		Time:     coder.EncodeTimeOpt(t),
		Key:      coder.EncodeData(key),
		Value:    coder.EncodeData(value),
	}

	return c.PostRaw(ctx, id, in)
}

func (c *Client) PostRaw(ctx context.Context, id logs.LogID, in PostIn) (int64, error) {
	var out PostOut
	err := c.H.Post(ctx, fmt.Sprintf("message/%s", id), in, &out)
	return out.NextOffset, err
}

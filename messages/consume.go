package messages

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/klev-dev/klev-api-go"
)

func (c *Client) Consume(ctx context.Context, id klev.LogID, opts ...klev.ConsumeOpt) (int64, []klev.ConsumeMessage, error) {
	copts := klev.ConsumeOpts{Encoding: klev.EncodingBase64}
	for _, opt := range opts {
		copts = opt(copts)
	}

	var out klev.ConsumeOut
	err := c.H.Get(ctx, fmt.Sprintf("messages/%s?%s", id, query(copts)), &out)
	if err != nil {
		return 0, nil, err
	}

	coder, err := klev.ParseMessageEncoding(out.Encoding)
	if err != nil {
		return 0, nil, err
	}

	var msgs = make([]klev.ConsumeMessage, len(out.Messages))
	for i, outMsg := range out.Messages {
		msg, err := outMsg.Decode(coder)
		if err != nil {
			return 0, nil, err
		}

		msgs[i] = msg
	}

	return out.NextOffset, msgs, err
}

func (c *Client) GetByOffset(ctx context.Context, id klev.LogID, offset int64) (klev.ConsumeMessage, error) {
	var out klev.GetOut
	err := c.H.Get(ctx, fmt.Sprintf("message/%s?offset=%d&encoding=base64", id, offset), &out)
	if err != nil {
		return klev.ConsumeMessage{}, err
	}

	return out.Decode()
}

func (c *Client) GetByKey(ctx context.Context, id klev.LogID, key []byte) (klev.ConsumeMessage, error) {
	coder := klev.EncodingBase64
	var out klev.GetOut
	err := c.H.Post(ctx, fmt.Sprintf("message/%s/key", id), klev.GetByKeyIn{
		Encoding: coder.String(),
		Key:      coder.EncodeData(key),
	}, &out)
	if err != nil {
		return klev.ConsumeMessage{}, err
	}

	return out.Decode()
}

func query(c klev.ConsumeOpts) string {
	var params []string
	if c.Offset != nil {
		params = append(params, fmt.Sprintf("offset=%d", *c.Offset))
	}
	if c.OffsetID != nil {
		params = append(params, fmt.Sprintf("offset_id=%s", *c.OffsetID))
	}
	if c.Size != nil {
		params = append(params, fmt.Sprintf("len=%d", *c.Size))
	}
	if c.Poll != nil {
		params = append(params, fmt.Sprintf("poll=%d", (*c.Poll)/time.Millisecond))
	}
	params = append(params, fmt.Sprintf("encoding=%s", c.Encoding))
	return strings.Join(params, "&")
}

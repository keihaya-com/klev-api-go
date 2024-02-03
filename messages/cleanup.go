package messages

import (
	"context"
	"fmt"
	"strings"

	"github.com/klev-dev/klev-api-go"
)

func (c *Client) Cleanup(ctx context.Context, id klev.LogID, opts ...klev.CleanupOpt) (int64, error) {
	var in klev.CleanupIn
	for _, opt := range opts {
		in = opt(in)
	}
	out, err := c.CleanupRaw(ctx, id, in)
	return out.Size, err
}

func (c *Client) CleanupRaw(ctx context.Context, id klev.LogID, in klev.CleanupIn) (klev.CleanupOut, error) {
	query := queryCleanup(in)

	var path = fmt.Sprintf("messages/%s", id)
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}

	var out klev.CleanupOut
	err := c.H.Delete(ctx, path, &out)
	return out, err
}

func queryCleanup(c klev.CleanupIn) string {
	var params []string
	if c.TrimSeconds > 0 {
		params = append(params, fmt.Sprintf("trim_seconds=%d", c.TrimSeconds))
	}
	if c.TrimSize > 0 {
		params = append(params, fmt.Sprintf("trim_size=%d", c.TrimSize))
	}
	if c.TrimCount > 0 {
		params = append(params, fmt.Sprintf("trim_count=%d", c.TrimCount))
	}
	if c.CompactSeconds > 0 {
		params = append(params, fmt.Sprintf("compact_seconds=%d", c.CompactSeconds))
	}
	if c.ExpireSeconds > 0 {
		params = append(params, fmt.Sprintf("expire_seconds=%d", c.ExpireSeconds))
	}
	return strings.Join(params, "&")
}

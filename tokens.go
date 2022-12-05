package api

import (
	"context"
	"fmt"
)

type TokenID string

type TokenIn struct {
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
}

type TokensOut struct {
	Tokens []TokenOut `json:"tokens,omitempty"`
}

type TokenOut struct {
	TokenID  TokenID  `json:"token_id"`
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
	Bearer   string   `json:"bearer,omitempty"`
}

func (c *Client) TokensList(ctx context.Context) ([]TokenOut, error) {
	var out TokensOut
	err := c.HTTPGet(ctx, fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (c *Client) TokenCreate(ctx context.Context, in TokenIn) (TokenOut, error) {
	var out TokenOut
	err := c.HTTPPost(ctx, fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (c *Client) TokenGet(ctx context.Context, id TokenID) (TokenOut, error) {
	var out TokenOut
	err := c.HTTPGet(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

func (c *Client) TokenDelete(ctx context.Context, id TokenID) error {
	var out TokenOut
	return c.HTTPDelete(ctx, fmt.Sprintf("token/%s", id), &out)
}

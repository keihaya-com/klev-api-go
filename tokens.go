package api

import (
	"context"
	"fmt"
)

type TokenID string

type TokenCreate struct {
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
}

type Tokens struct {
	Tokens []Token `json:"tokens,omitempty"`
}

type Token struct {
	TokenID  TokenID  `json:"token_id"`
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
	Bearer   string   `json:"bearer,omitempty"`
}

func (c *Client) TokensList(ctx context.Context) ([]Token, error) {
	var out Tokens
	err := c.httpGet(ctx, fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (c *Client) TokensFind(ctx context.Context, metadata string) ([]Token, error) {
	var out Tokens
	err := c.httpGet(ctx, fmt.Sprintf("tokens?q=%s", metadata), &out)
	return out.Tokens, err
}

func (c *Client) TokenCreate(ctx context.Context, in TokenCreate) (Token, error) {
	var out Token
	err := c.httpPost(ctx, fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (c *Client) TokenGet(ctx context.Context, id TokenID) (Token, error) {
	var out Token
	err := c.httpGet(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

func (c *Client) TokenDelete(ctx context.Context, id TokenID) (Token, error) {
	var out Token
	err := c.httpDelete(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

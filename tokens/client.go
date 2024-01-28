package tokens

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go/client"
)

type TokenID string

type Token struct {
	TokenID  TokenID  `json:"token_id"`
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
	Bearer   string   `json:"bearer,omitempty"`
}

type Tokens struct {
	Tokens []Token `json:"tokens,omitempty"`
}

type CreateParams struct {
	Metadata string   `json:"metadata"`
	ACL      []string `json:"acl"`
}

type Client struct {
	H client.HTTP
}

func (c *Client) List(ctx context.Context) ([]Token, error) {
	var out Tokens
	err := c.H.Get(ctx, fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]Token, error) {
	var out Tokens
	err := c.H.Get(ctx, fmt.Sprintf("tokens?metadata=%s", metadata), &out)
	return out.Tokens, err
}

func (c *Client) Create(ctx context.Context, in CreateParams) (Token, error) {
	var out Token
	err := c.H.Post(ctx, fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id TokenID) (Token, error) {
	var out Token
	err := c.H.Get(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id TokenID) (Token, error) {
	var out Token
	err := c.H.Delete(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

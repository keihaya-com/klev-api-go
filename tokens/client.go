package tokens

import (
	"context"
	"fmt"

	"github.com/klev-dev/klev-api-go"
)

type Client struct {
	H klev.HTTP
}

func New(cfg klev.Config) *Client {
	return &Client{klev.New(cfg)}
}

func (c *Client) List(ctx context.Context) ([]klev.Token, error) {
	var out klev.Tokens
	err := c.H.Get(ctx, fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (c *Client) Find(ctx context.Context, metadata string) ([]klev.Token, error) {
	var out klev.Tokens
	err := c.H.Get(ctx, fmt.Sprintf("tokens?metadata=%s", metadata), &out)
	return out.Tokens, err
}

func (c *Client) Create(ctx context.Context, in klev.TokenCreateParams) (klev.Token, error) {
	var out klev.Token
	err := c.H.Post(ctx, fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (c *Client) Get(ctx context.Context, id klev.TokenID) (klev.Token, error) {
	var out klev.Token
	err := c.H.Get(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

func (c *Client) Delete(ctx context.Context, id klev.TokenID) (klev.Token, error) {
	var out klev.Token
	err := c.H.Delete(ctx, fmt.Sprintf("token/%s", id), &out)
	return out, err
}

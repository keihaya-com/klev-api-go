package api

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

type TokenIn struct {
	Metadata string `json:"metadata"`
}

type TokensOut struct {
	Tokens []TokenOut `json:"tokens,omitempty"`
}

type TokenOut struct {
	TokenID  ksuid.KSUID `json:"token_id"`
	Metadata string      `json:"metadata"`
	Bearer   string      `json:"bearer,omitempty"`
}

func (c *Client) TokensList() ([]TokenOut, error) {
	var out TokensOut
	err := c.HTTPGet(fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (c *Client) TokenCreate(in TokenIn) (TokenOut, error) {
	var out TokenOut
	err := c.HTTPPost(fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (c *Client) TokenGet(tokenID ksuid.KSUID) error {
	var out TokenOut
	return c.HTTPGet(fmt.Sprintf("token/%s", tokenID), &out)
}

func (c *Client) TokenDelete(tokenID ksuid.KSUID) error {
	var out TokenOut
	return c.HTTPDelete(fmt.Sprintf("token/%s", tokenID), &out)
}

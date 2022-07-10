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

func (k *Client) TokensList() ([]TokenOut, error) {
	var out TokensOut
	err := k.base.Get(fmt.Sprintf("tokens"), &out)
	return out.Tokens, err
}

func (k *Client) TokenCreate(in TokenIn) (TokenOut, error) {
	var out TokenOut
	err := k.base.Post(fmt.Sprintf("tokens"), in, &out)
	return out, err
}

func (k *Client) TokenGet(tokenID ksuid.KSUID) error {
	var out TokenOut
	return k.base.Get(fmt.Sprintf("token/%s", tokenID), &out)
}

func (k *Client) TokenDelete(tokenID ksuid.KSUID) error {
	var out TokenOut
	return k.base.Delete(fmt.Sprintf("token/%s", tokenID), &out)
}

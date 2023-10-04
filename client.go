package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Config is used to configure a client instance
type Config struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

// NewConfig creates new client configuration with a token
func NewConfig(token string) Config {
	return Config{
		BaseURL: "https://api.klev.dev",
		Token:   token,
		Client:  http.DefaultClient,
	}
}

// Client wraps interactions with klev api
type Client struct {
	baseURL string
	token   string
	client  *http.Client
}

// New create a new client from a config
func New(cfg Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
		client:  cfg.Client,
	}
}

func (c *Client) Paths(ctx context.Context) (map[string]string, error) {
	var out map[string]string
	err := c.httpGet(ctx, "", &out)
	return out, err
}

func (c *Client) httpGet(ctx context.Context, path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP get: %w", err)
	}

	return c.httpDo(req, out)
}

func (c *Client) httpPost(ctx context.Context, path string, in interface{}, out interface{}) error {
	return c.httpUpdate(ctx, http.MethodPost, path, in, out)
}

func (c *Client) httpPatch(ctx context.Context, path string, in interface{}, out interface{}) error {
	return c.httpUpdate(ctx, http.MethodPatch, path, in, out)
}

func (c *Client) httpUpdate(ctx context.Context, method string, path string, in interface{}, out interface{}) error {
	bin, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewBuffer(bin))
	if err != nil {
		return fmt.Errorf("could not prepare HTTP post: %w", err)
	}

	return c.httpDo(req, out)
}

func (c *Client) httpDelete(ctx context.Context, path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP delete: %w", err)
	}

	return c.httpDo(req, out)
}

func (c *Client) httpDo(req *http.Request, out interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not execute request: %w", err)
	}

	defer resp.Body.Close()
	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var eout ErrorOut
		err = json.Unmarshal(bout, &eout)
		if err != nil {
			return fmt.Errorf("could not unmarshal error: %w", err)
		}
		return &eout
	}

	err = json.Unmarshal(bout, out)
	if err != nil {
		return fmt.Errorf("could not unmarshal response: %w", err)
	}
	return nil
}

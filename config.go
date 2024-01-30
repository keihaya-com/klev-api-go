package klev

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

type HTTP interface {
	Get(ctx context.Context, path string, out any) error
	Post(ctx context.Context, path string, in any, out any) error
	Patch(ctx context.Context, path string, in any, out any) error
	Delete(ctx context.Context, path string, out any) error
}

func New(cfg Config) HTTP {
	return &httpClient{
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
		client:  cfg.Client,
	}
}

type httpClient struct {
	baseURL string
	token   string
	client  *http.Client
}

func (c *httpClient) Get(ctx context.Context, path string, out any) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP get: %w", err)
	}

	return c.do(req, out)
}

func (c *httpClient) Post(ctx context.Context, path string, in any, out any) error {
	return c.update(ctx, http.MethodPost, path, in, out)
}

func (c *httpClient) Patch(ctx context.Context, path string, in any, out any) error {
	return c.update(ctx, http.MethodPatch, path, in, out)
}

func (c *httpClient) update(ctx context.Context, method string, path string, in any, out any) error {
	bin, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, method, target, bytes.NewBuffer(bin))
	if err != nil {
		return fmt.Errorf("could not prepare HTTP post: %w", err)
	}

	return c.do(req, out)
}

func (c *httpClient) Delete(ctx context.Context, path string, out any) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP delete: %w", err)
	}

	return c.do(req, out)
}

func (c *httpClient) do(req *http.Request, out any) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("could not execute request: %w", err)
	}

	defer resp.Body.Close()
	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}

	if resp.StatusCode >= 400 {
		var eout APIError
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

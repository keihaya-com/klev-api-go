package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Config struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func NewConfig(token string) Config {
	return Config{
		BaseURL: "https://api.klev.dev",
		Token:   token,
		Client:  http.DefaultClient,
	}
}

type Client struct {
	baseURL string
	token   string
	client  *http.Client
}

func New(cfg Config) *Client {
	return &Client{
		baseURL: cfg.BaseURL,
		token:   cfg.Token,
		client:  cfg.Client,
	}
}

func (c *Client) HTTPGet(path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP get: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return c.HTTPDo(req, out)
}

func (c *Client) HTTPPost(path string, in interface{}, out interface{}) error {
	bin, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(bin))
	if err != nil {
		return fmt.Errorf("could not prepare HTTP post: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return c.HTTPDo(req, out)
}

func (c *Client) HTTPDelete(path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", c.baseURL, path)

	req, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP delete: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return c.HTTPDo(req, out)
}

func (c *Client) HTTPDo(req *http.Request, out interface{}) error {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not execute request: %w", err)
	}

	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}
	resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(bout, out)
		if err != nil {
			return fmt.Errorf("could not unmarshal: %w", err)
		}
		return nil
	} else {
		var eout ErrorOut
		err = json.Unmarshal(bout, &eout)
		if err != nil {
			return fmt.Errorf("could not unmarshal error: %w", err)
		}
		return fmt.Errorf("(%s) %s", eout.Code, eout.Message)
	}
}

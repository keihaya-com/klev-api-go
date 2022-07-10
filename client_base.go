package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

type baseClient struct {
	baseURL string
	token   string
	log     *zap.SugaredLogger
}

func newBase(baseURL string, log *zap.SugaredLogger) baseClient {
	token := os.Getenv("KLEV_TOKEN")
	if token == "" {
		panic("fish> set -x KLEV_TOKEN '$'\nbash> ?export")
	}

	return baseClient{
		baseURL: baseURL,
		token:   token,
		log:     log,
	}
}

func (t baseClient) Get(path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", t.baseURL, path)

	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP get: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return t.request(req, out)
}

func (t baseClient) Post(path string, in interface{}, out interface{}) error {
	bin, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("could not marshal json: %w", err)
	}

	target := fmt.Sprintf("%s/%s", t.baseURL, path)

	req, err := http.NewRequest(http.MethodPost, target, bytes.NewBuffer(bin))
	if err != nil {
		return fmt.Errorf("could not prepare HTTP post: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return t.request(req, out)
}

func (t baseClient) Delete(path string, out interface{}) error {
	target := fmt.Sprintf("%s/%s", t.baseURL, path)

	req, err := http.NewRequest(http.MethodDelete, target, nil)
	if err != nil {
		return fmt.Errorf("could not prepare HTTP delete: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return t.request(req, out)
}

func (t baseClient) request(req *http.Request, out interface{}) error {
	t.log.Debugw("-> request", "url", req.URL, "method", req.Method)
	defer func(start time.Time) {
		t.log.Debugw("<- request", "url", req.URL, "method", req.Method, "t", time.Since(start))
	}(time.Now())

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.token))
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

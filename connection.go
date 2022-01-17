package oanda

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"path"
	"time"

	"github.com/pkg/errors"
)

type OandaEnvironment int

const (
	OandaPractice OandaEnvironment = iota + 1
	OandaLive
	oandaDummy
)

type Connection struct {
	Token       string
	Environemnt OandaEnvironment
	Timeout     time.Duration
	Strict      bool
}

func (c *Connection) request(ctx context.Context, params *requestParams) (*http.Response, error) {
	destURL := oandaBaseURL(c.Environemnt).rest
	destURL.Path = path.Join(destURL.Path, params.endPoint)

	var reader io.Reader
	if params.body != nil {
		body, _ := json.Marshal(params.body)
		reader = bytes.NewBuffer(body)
	}

	req, err := http.NewRequestWithContext(ctx, params.method, destURL.String(), reader)
	if err != nil {
		return nil, errors.Errorf("Prepare new request failed: %v", err)
	}

	// req.Header.Set("User-Agent", "Go 1.1 package http")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	for _, h := range params.headers {
		req.Header.Set(h.key, h.value)
	}

	reqQ := req.URL.Query()
	for _, q := range params.queries {
		reqQ.Add(q.key, q.value)
	}
	req.URL.RawQuery = reqQ.Encode()

	client := http.Client{
		Timeout: c.Timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Errorf("Request canceled: %v", err)
	}

	return resp, nil
}

func (c *Connection) stream(ctx context.Context, params *requestParams) (*http.Response, error) {
	destURL := oandaBaseURL(c.Environemnt).stream
	destURL.Path = path.Join(destURL.Path, params.endPoint)

	req, err := http.NewRequestWithContext(ctx, params.method, destURL.String(), nil)
	if err != nil {
		return nil, errors.Errorf("error in stream method: %v", err)
	}

	// req.Header.Set("User-Agent", "Go 1.1 package http")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", "application/json")
	for _, h := range params.headers {
		req.Header.Set(h.key, h.value)
	}

	reqQ := req.URL.Query()
	for _, q := range params.queries {
		reqQ.Add(q.key, q.value)
	}
	req.URL.RawQuery = reqQ.Encode()

	client := http.Client{
		Timeout: 0,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Errorf("error in stream method: %v", err)
	}

	return resp, nil
}

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Headers    http.Header
	APIKey     string
}

type RequestConfig struct {
	Method  string
	Path    string
	Query   map[string]string
	Body    interface{}
	Headers map[string]string
}

func NewClient(baseUrl string, timeout time.Duration) (*Client, error) {
	parsedUrl, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseURL: parsedUrl,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
		Headers: make(http.Header),
	}, nil
}

func (c *Client) NewRequest(ctx context.Context, config RequestConfig) (*http.Request, error) {
	rel := &url.URL{Path: config.Path}
	u := c.BaseURL.ResolveReference(rel)

	query := u.Query()
	for k, v := range config.Query {
		query.Set(k, v)
	}
	u.RawQuery = query.Encode()

	var bodyReader *bytes.Reader
	if config.Body != nil {
		jsonBody, err := json.Marshal(config.Body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	var req *http.Request
	var err error
	if bodyReader != nil {
		req, err = http.NewRequestWithContext(ctx, config.Method, u.String(), bodyReader)
	} else {
		req, err = http.NewRequestWithContext(ctx, config.Method, u.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// Headers globali
	for k, values := range c.Headers {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}

	// Headers specifici per request
	for k, v := range config.Headers {
		req.Header.Set(k, v)
	}

	if config.Body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

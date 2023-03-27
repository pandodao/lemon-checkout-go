package lemon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const ApiBase = "https://api.lemonsqueezy.com/v1"

type Client struct {
	key    string
	debug  bool
	logger Logger
}

type Logger interface {
	Debugf(format string, args ...interface{})
}

type Option func(*Client)

func WithDebug(debug bool) Option {
	return func(c *Client) {
		c.debug = debug
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func New(key string, opts ...Option) *Client {
	c := &Client{
		key: key,
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Client) debugLog(v string, args ...interface{}) {
	if !c.debug {
		return
	}
	if c.logger != nil {
		c.logger.Debugf(v, args...)
	} else {
		log.Printf(v, args...)
	}
}

func (c *Client) request(ctx context.Context, method string, uri string, body, result any) error {
	reqLog := fmt.Sprintf("[Request] %s %s", method, uri)
	start := time.Now()
	var r io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		reqLog += fmt.Sprintf(" %s", string(data))
		r = bytes.NewBuffer(data)
	}
	c.debugLog(reqLog)

	req, err := http.NewRequestWithContext(ctx, method, ApiBase+uri, r)
	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/vnd.api+json")
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("Authorization", "Bearer "+c.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusCreated {
		return http.ErrNotSupported
	}

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.debugLog("[Response %s] %s %s %s", time.Since(start), method, uri, string(respData))

	var res GeneralWrapper

	res.Data = result

	if err := json.Unmarshal(respData, &res); err != nil {
		return err
	}

	return nil
}

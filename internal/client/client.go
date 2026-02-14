package client

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Softorize/bwm/internal/models"
)

const baseURL = "https://ssl.bing.com/webmaster/api.svc/json"

type Client struct {
	apiKey     string
	httpClient *http.Client
	maxRetries int
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		maxRetries: 3,
	}
}

func (c *Client) get(method string, params url.Values, result any) error {
	u := fmt.Sprintf("%s/%s", baseURL, method)
	if params == nil {
		params = url.Values{}
	}
	params.Set("apikey", c.apiKey)
	fullURL := u + "?" + params.Encode()

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			time.Sleep(backoff)
		}

		resp, err := c.httpClient.Get(fullURL)
		if err != nil {
			lastErr = err
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}

		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("server error: %d %s", resp.StatusCode, string(body))
			continue
		}

		if resp.StatusCode >= 400 {
			var apiErr models.APIError
			if json.Unmarshal(body, &apiErr) == nil && apiErr.Message != "" {
				return &apiErr
			}
			return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
		}

		return unwrap(body, result)
	}
	return fmt.Errorf("after %d retries: %w", c.maxRetries, lastErr)
}

func (c *Client) post(method string, payload any, result any) error {
	u := fmt.Sprintf("%s/%s?apikey=%s", baseURL, method, url.QueryEscape(c.apiKey))

	var bodyReader io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("marshaling request: %w", err)
		}
		bodyReader = strings.NewReader(string(data))
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := time.Duration(math.Pow(2, float64(attempt-1))) * time.Second
			time.Sleep(backoff)
		}

		var resp *http.Response
		var err error
		if bodyReader != nil {
			data, _ := json.Marshal(payload)
			resp, err = c.httpClient.Post(u, "application/json", strings.NewReader(string(data)))
		} else {
			resp, err = c.httpClient.Post(u, "application/json", nil)
		}
		if err != nil {
			lastErr = err
			continue
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			lastErr = err
			continue
		}

		if resp.StatusCode >= 500 {
			lastErr = fmt.Errorf("server error: %d %s", resp.StatusCode, string(body))
			continue
		}

		if resp.StatusCode >= 400 {
			var apiErr models.APIError
			if json.Unmarshal(body, &apiErr) == nil && apiErr.Message != "" {
				return &apiErr
			}
			return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
		}

		if result == nil {
			return nil
		}
		return unwrap(body, result)
	}
	return fmt.Errorf("after %d retries: %w", c.maxRetries, lastErr)
}

// unwrap removes the Bing "d" envelope: {"d": <actual data>}
func unwrap(body []byte, result any) error {
	var envelope struct {
		D json.RawMessage `json:"d"`
	}
	if err := json.Unmarshal(body, &envelope); err == nil && envelope.D != nil {
		return json.Unmarshal(envelope.D, result)
	}
	// Fallback: no envelope
	return json.Unmarshal(body, result)
}

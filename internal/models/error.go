package models

import "fmt"

type APIError struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Bing API error %d: %s", e.ErrorCode, e.Message)
}

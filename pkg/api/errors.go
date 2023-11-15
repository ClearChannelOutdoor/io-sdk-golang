package api

import (
	"fmt"
	"time"
)

type APIError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Title   string `json:"title"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s (%d): %s", e.Title, e.Status, e.Message)
}

type retryError struct {
	Err        error
	RetryAfter time.Duration
}

func (e retryError) Error() string {
	return fmt.Sprintf("retry after %s: %s", e.RetryAfter, e.Err)
}

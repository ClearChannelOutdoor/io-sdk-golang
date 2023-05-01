package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

const (
	authorizationHeader     string = "Authorization"
	bearerFmt               string = "Bearer %s"
	defaultMaxAttempts      uint   = 5
	defaultThrottleRetrySec uint   = 5
)

type apiError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Title   string `json:"title"`
}

func (e apiError) Error() string {
	return fmt.Sprintf("%s (%d): %s", e.Title, e.Status, e.Message)
}

type retryError struct {
	Err        error
	RetryAfter time.Duration
}

func (e retryError) Error() string {
	return fmt.Sprintf("retry after %s: %s", e.RetryAfter, e.Err)
}

type Endpoint struct {
	c           *http.Client
	BearerToken string
	Environment string
	Host        string
	MaxAttempts uint
	Proto       string
}

func NewEndpoint(tkn, env, host, proto string) *Endpoint {
	return &Endpoint{
		c:           &http.Client{},
		BearerToken: tkn,
		Environment: env,
		Host:        host,
		MaxAttempts: defaultMaxAttempts,
		Proto:       proto,
	}
}

func (e *Endpoint) retry(method string, reqPath string, opts ...Options) (map[string]any, error) {
	// determine the URL
	url := fmt.Sprintf("%s://%s%s", e.Proto, e.Host, reqPath)

	// build the request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// set the bearer token
	req.Header.Set(authorizationHeader, fmt.Sprintf(bearerFmt, e.BearerToken))

	// set the query params as appropriate
	if len(opts) > 0 {
		req.URL.RawQuery = opts[0].Query().Encode()
	}

	// track what is actually provided back from the API
	dataMap := map[string]any{}

	// create retry operation
	apiRequest := func() error {
		res, err := e.c.Do(req)
		if err != nil {
			// retry on request errors (network issues are caught here)
			return &retryError{
				Err: err,
			}
		}

		// process the response
		defer res.Body.Close()

		// read the body
		bdy, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// check to see if the response contains an error status
		if res.StatusCode > 399 {
			// attempt to unmarshal the body into an apiError
			var apiErr apiError
			if err := json.Unmarshal(bdy, &apiErr); err != nil {
				// the error is not an apiError, it is a string value
				apiErr = apiError{
					Message: string(bdy),
					Status:  res.StatusCode,
					Title:   res.Status,
				}
			}

			// check for throttling...
			if apiErr.Status == http.StatusTooManyRequests {
				// set RetryAfter to delay the next request for throttling
				return &retryError{
					Err:        apiErr,
					RetryAfter: time.Second * time.Duration(defaultThrottleRetrySec),
				}
			}

			// when the error is a client created error, do not retry
			if res.StatusCode < http.StatusInternalServerError {
				return apiErr
			}

			// otherwise retry
			return &retryError{
				Err: apiErr,
			}
		}

		// unmarshal the body
		var data interface{}
		if err := json.Unmarshal(bdy, &data); err != nil {
			// check if data is a string type
			if _, ok := data.(string); ok {
				data = map[string]string{
					"message": data.(string),
				}
			} else {
				// return unmarshal error
				return err
			}
		}

		// reset the data map to whatever was unmarshaled
		dataMap = data.(map[string]any)

		// return no error
		return nil
	}

	// make the request with retry logic
	if err := retry.Do(
		apiRequest,
		retry.Attempts(e.MaxAttempts),
		retry.RetryIf(func(err error) bool {
			// retry on retry errors
			if _, ok := err.(*retryError); ok {
				return true
			}

			return false
		}),
		retry.DelayType(func(n uint, err error, config *retry.Config) time.Duration {
			// check if the error is a retry error
			if retryErr, ok := err.(*retryError); ok {
				// return the retry delay if provided (may be empty)
				if retryErr.RetryAfter > 0 {
					return retryErr.RetryAfter
				}
			}

			// return the default retry delay
			return retry.BackOffDelay(n, err, config)
		}),
	); err != nil {
		return nil, err
	}

	return dataMap, nil
}

func (e *Endpoint) Request(method string, path string, opts ...Options) (map[string]any, error) {
	r, err := e.retry(method, path, opts...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

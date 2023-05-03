package api

import (
	"bytes"
	"context"
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
	deleteMethod            string = "DELETE"
	getMethod               string = "GET"
	patchMethod             string = "PATCH"
	postMethod              string = "POST"
	putMethod               string = "PUT"
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

type Endpoint[T any] struct {
	c           *http.Client
	BearerToken string
	Context     context.Context
	Environment string
	Host        string
	MaxAttempts uint
	Path        string
	Proto       string
}

func NewEndpoint[T any](env Environment, path string) *Endpoint[T] {
	return &Endpoint[T]{
		c:           &http.Client{},
		BearerToken: env.Token,
		Context:     retry.DefaultContext,
		Environment: env.Name,
		Host:        env.Host,
		MaxAttempts: defaultMaxAttempts,
		Path:        path,
		Proto:       env.Proto,
	}
}

func (e *Endpoint[T]) retry(method string, reqPath string, body io.Reader, opts ...Options) ([]byte, error) {
	// determine the URL
	url := fmt.Sprintf("%s://%s%s", e.Proto, e.Host, reqPath)

	// build the request
	req, err := http.NewRequestWithContext(e.Context, method, url, body)
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
	data := []byte{}

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

		// assign the body
		data = bdy

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

	return data, nil
}

func (e *Endpoint[T]) request(method string, path string, body io.Reader, opts ...Options) ([]byte, error) {
	r, err := e.retry(method, path, body, opts...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (e *Endpoint[T]) Create(mdl T) (T, error) {
	var created T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return created, err
	}

	// make the request to the API
	data, err := e.request(postMethod, e.Path, bytes.NewBuffer(jd))
	if err != nil {
		return created, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &created); err != nil {
		return created, err
	}

	return created, nil
}

func (e *Endpoint[T]) Delete(id string) error {
	// make the request to the API
	_, err := e.request(deleteMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return err
	}

	return nil
}

func (e *Endpoint[T]) Get(id string) (T, error) {
	var mdl T

	// make the request to the API
	data, err := e.request(getMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return mdl, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &mdl); err != nil {
		return mdl, err
	}

	return mdl, nil
}

func (e *Endpoint[T]) GetAll(opts ...Options) (SearchResult[T], error) {
	var sr SearchResult[T]

	// make the request to the API
	data, err := e.request(getMethod, e.Path, nil, opts...)
	if err != nil {
		return sr, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &sr); err != nil {
		return sr, err
	}

	return sr, nil
}

func (e *Endpoint[T]) Patch(id string, mdl T) (T, error) {
	var patched T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return patched, err
	}

	// make the request to the API
	data, err := e.request(patchMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return patched, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &patched); err != nil {
		return patched, err
	}

	return patched, nil
}

func (e *Endpoint[T]) Update(id string, mdl T) (T, error) {
	var updated T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return updated, err
	}

	// make the request to the API
	data, err := e.request(putMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return updated, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &updated); err != nil {
		return updated, err
	}

	return updated, nil
}

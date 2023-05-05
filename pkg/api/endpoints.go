package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"time"

	"github.com/avast/retry-go"
	"golang.org/x/oauth2"
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

var retryAfterRE *regexp.Regexp = regexp.MustCompile(`retry after (\d+)s\: `)

type Endpoint[T any] struct {
	c           *http.Client
	ctx         context.Context
	svc         *Service
	MaxAttempts uint
	OAuthToken  *oauth2.Token
	Path        string
}

func NewEndpoint[T any](svc *Service, path string) *Endpoint[T] {
	return &Endpoint[T]{
		c:           &http.Client{},
		svc:         svc,
		ctx:         retry.DefaultContext,
		MaxAttempts: defaultMaxAttempts,
		Path:        path,
	}
}

func (e *Endpoint[T]) ensureBearerToken() (string, error) {
	// if the token is blank or expired, get a new one
	if e.OAuthToken == nil || (!e.OAuthToken.Expiry.IsZero() && e.OAuthToken.Expiry.Before(time.Now())) {
		// set the auth style to header
		e.svc.oauth2.AuthStyle = oauth2.AuthStyleInHeader

		// retrieve the token
		tkn, err := e.svc.oauth2.Token(e.ctx)
		if err != nil {
			return "", err
		}

		// assign the token
		e.OAuthToken = tkn
	}

	// return the access token
	return e.OAuthToken.AccessToken, nil
}

func (e *Endpoint[T]) retry(method string, reqPath string, body io.Reader, opts ...*Options) ([]byte, error) {
	// determine the URL
	url := fmt.Sprintf("%s://%s%s", e.svc.Proto, e.svc.Host, reqPath)

	// build the request
	req, err := http.NewRequestWithContext(e.ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	// set the bearer token authorization header
	tkn, err := e.ensureBearerToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set(authorizationHeader, fmt.Sprintf(bearerFmt, tkn))

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
		switch e := err.(type) {
		// when the error is a wrapped error, clean it up and return the last error
		case retry.Error:
			el := e.WrappedErrors()
			err = el[len(el)-1]
			if retryAfterRE.MatchString(err.Error()) {
				i := retryAfterRE.FindStringIndex(err.Error())
				err = errors.New(err.Error()[i[1]:])
			}
		}

		return nil, err
	}

	return data, nil
}

func (e *Endpoint[T]) request(method string, path string, body io.Reader, opts ...*Options) ([]byte, error) {
	r, err := e.retry(method, path, body, opts...)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Create creates a new model of type T within the API
func (e *Endpoint[T]) Create(mdl *T) error {
	var created T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(postMethod, e.Path, bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &created); err != nil {
		return err
	}

	// assign the created model
	*mdl = created

	return nil
}

// Delete deletes a model of type T from the API
func (e *Endpoint[T]) Delete(id string) error {
	// make the request to the API
	_, err := e.request(deleteMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return err
	}

	return nil
}

// Get gets a single model of type T from the API
func (e *Endpoint[T]) Get(id string) (*T, error) {
	var mdl T

	// make the request to the API
	data, err := e.request(getMethod, fmt.Sprintf("%s/%s", e.Path, id), nil)
	if err != nil {
		return nil, err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &mdl); err != nil {
		return nil, err
	}

	return &mdl, nil
}

// Patch updates a model of type T within the API while preserving
// any fields that are omitted and any additional externalIDs
// that already exist
func (e *Endpoint[T]) Patch(id string, mdl *T) error {
	var patched T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(patchMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &patched); err != nil {
		return err
	}

	// assign the patched model
	*mdl = patched

	return nil
}

// Search searches for models of type T within the API
func (e *Endpoint[T]) Search(opts ...*Options) (SearchResult[T], error) {
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

// Update updates a model of type T within the API
func (e *Endpoint[T]) Update(id string, mdl *T) error {
	var updated T

	jd, err := json.Marshal(mdl)
	if err != nil {
		return err
	}

	// make the request to the API
	data, err := e.request(putMethod, fmt.Sprintf("%s/%s", e.Path, id), bytes.NewBuffer(jd))
	if err != nil {
		return err
	}

	// unmarshal the data into the struct
	if err := json.Unmarshal(data, &updated); err != nil {
		return err
	}

	// assign the updated model
	*mdl = updated

	return nil
}

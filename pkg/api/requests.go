package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/avast/retry-go"
	"golang.org/x/oauth2"
)

const (
	defaultMaxAttempts            uint          = 5
	defaultTokenExpirationOverlap time.Duration = 5 * time.Minute
)

func ensureBearerToken(ctx context.Context, a *api) (string, error) {
	if a.Svc == nil || a.Svc.oauth2 == nil {
		return "", nil
	}

	n := time.Now()

	// if the token is blank or expired, get a new one
	// removing check for OAuthToken.Expiry.Zero - this would indicate token does not expire
	a.Mu.Lock()
	defer a.Mu.Unlock()
	if a.OAuthToken == nil || a.OAuthToken.Expiry.Before(n.Add(-defaultTokenExpirationOverlap)) {
		// set the auth style to header
		a.Svc.oauth2.AuthStyle = oauth2.AuthStyleInHeader

		// retrieve the token
		tkn, err := a.Svc.oauth2.Token(ctx)
		if err != nil {
			return "", err
		}

		// assign the token
		a.OAuthToken = tkn
	}

	// return the access token
	return a.OAuthToken.AccessToken, nil
}

func retryRequest(ctx context.Context, hdr *http.Header, a *api, method string, reqPath string, body io.Reader, opts ...*Options) ([]byte, int, error) {
	// determine the URL
	url := fmt.Sprintf("%s://%s%s", a.Svc.Proto, a.Svc.Host, reqPath)

	// build the request
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, 0, err
	}

	// set the bearer token authorization header
	tkn, err := ensureBearerToken(ctx, a)
	if err != nil {
		return nil, http.StatusUnauthorized, err
	}
	req.Header.Set(authorizationHeader, fmt.Sprintf(bearerFmt, tkn))

	// set the headers (if applicable)
	if hdr != nil {
		for k, v := range *hdr {
			req.Header.Set(k, strings.Join(v, ","))
		}
	}

	// set the query params as appropriate
	if len(opts) > 0 {
		req.URL.RawQuery = opts[0].FormatQuery()
	}

	// track what is actually provided back from the API
	code := http.StatusProcessing
	data := []byte{}

	// create retry operation
	apiRequest := func() error {
		res, err := a.Clnt.Do(req)
		if err != nil {
			// retry on request errors (network issues are caught here)
			return &retryError{
				Err: err,
			}
		}

		// process the response
		defer res.Body.Close()

		// assign the response status code
		code = res.StatusCode

		// read the body
		bdy, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		// check to see if the response contains an error status
		if code > 399 {
			// attempt to unmarshal the body into an apiError
			var apiErr APIError
			if err := json.Unmarshal(bdy, &apiErr); err != nil {
				// the error is not an apiError, it is a string value
				apiErr = APIError{
					Message: string(bdy),
					Status:  code,
					Title:   res.Status,
				}
			}

			// check for throttling...
			if code == http.StatusTooManyRequests {
				// set RetryAfter to delay the next request for throttling
				return &retryError{
					Err:        apiErr,
					RetryAfter: time.Second * time.Duration(defaultThrottleRetrySec),
				}
			}

			// when the error is a client created error, do not retry
			if code < http.StatusInternalServerError {
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
		retry.Attempts(defaultMaxAttempts),
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
		if re, ok := err.(retry.Error); ok {
			el := re.WrappedErrors()

			err = el[0]
			if retryAfterRE.MatchString(err.Error()) {
				i := retryAfterRE.FindStringIndex(err.Error())
				err = errors.New(err.Error()[i[1]:])
			}
		}

		return nil, code, err
	}

	return data, code, nil
}

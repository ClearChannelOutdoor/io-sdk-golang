package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func Test_ensureBearerToken(t *testing.T) {
	var ts *httptest.Server
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	type fields struct {
		oauth2 *clientcredentials.Config
		token  *oauth2.Token
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"should properly get a bearer token",
			fields{
				oauth2: &clientcredentials.Config{
					ClientID:     "test-client-id",
					ClientSecret: "test-client-secret",
					TokenURL:     "test-token-url",
				},
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					io.WriteString(w, fmt.Sprintf(`{"access_token": "%s", "expires_in": 86400, "token_type": "bearer"}`, defaultTestBearerToken))
				},
			},
			defaultTestBearerToken,
			false,
		},
		{
			"should properly use existing bearer token",
			fields{
				oauth2: &clientcredentials.Config{
					ClientID:     "test-client-id",
					ClientSecret: "test-client-secret",
					TokenURL:     "test-token-url",
				},
				token: &oauth2.Token{
					AccessToken: defaultTestBearerToken,
					TokenType:   "bearer",
				},
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")
					io.WriteString(w, fmt.Sprintf(`{"access_token": "%s", "expires_in": 86400, "token_type": "bearer"}`, defaultTestBearerToken))
				},
			},
			defaultTestBearerToken,
			false,
		},
	}
	for _, tt := range tests {
		// reset the test server
		ts = getTestServer(tt.args.res)
		defer ts.Close()

		u, _ := url.Parse(ts.URL)
		t.Run(tt.name, func(t *testing.T) {
			e := NewEndpoint[TestModel](
				&Service{
					oauth2: tt.fields.oauth2,
					Host:   u.Host,
					Proto:  u.Scheme,
				},
				"/test",
			)

			// set a token url pointing back to our test server
			e.a.Svc.oauth2.TokenURL = fmt.Sprintf("%s/v2/token", ts.URL)
			e.a.OAuthToken = tt.fields.token

			got, err := ensureBearerToken(context.Background(), e.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.ensureBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Endpoint.ensureBearerToken() = \n\t%s\n want \n\t%s", got, tt.want)
			}
		})
	}
}

func Test_retryRequest(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}
	retryCount := 0

	testData, _ := json.Marshal(map[string]any{
		"req.URL": "/test",
		"test":    "test",
	})

	var ts *httptest.Server

	type args struct {
		res    func(w http.ResponseWriter, r *http.Request)
		method string
		path   string
		body   io.Reader
		opts   []*Options
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			"should properly succeed in issuing a request",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					data := map[string]any{}
					data["test"] = "test"
					data["req.URL"] = r.URL.String()
					_ = json.NewEncoder(w).Encode(data)
				},
				"GET",
				"/test",
				nil,
				[]*Options{},
			},
			testData,
			false,
		},
		{
			"should properly retry on failures that are 500 or above",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					if retryCount < 4 {
						retryCount++
						w.WriteHeader(http.StatusInternalServerError)
						w.Header().Set("Content-Type", "application/json")

						data := map[string]any{}
						data["message"] = "test error"
						data["status"] = 500
						data["title"] = "Internal Server Error"
						_ = json.NewEncoder(w).Encode(data)
						return
					}

					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					data := map[string]any{}
					data["test"] = "test"
					data["req.URL"] = r.URL.String()
					_ = json.NewEncoder(w).Encode(data)
				},
				"GET",
				"/test",
				nil,
				[]*Options{},
			},
			testData,
			false,
		},
		{
			"should properly retry on too many requests error",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					if retryCount < 1 {
						retryCount++

						w.WriteHeader(http.StatusTooManyRequests)
						return
					}

					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					data := map[string]any{}
					data["test"] = "test"
					data["req.URL"] = r.URL.String()
					_ = json.NewEncoder(w).Encode(data)
				},
				"GET",
				"/test",
				nil,
				[]*Options{},
			},
			testData,
			false,
		},
		{
			"should not retry on 4xx errors",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")
				},
				"GET",
				"/test",
				nil,
				[]*Options{},
			},
			nil,
			true,
		},
		{
			"should properly handle a string (non-JSON) response",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/text")
					w.Write([]byte("a random string"))
				},
				"GET",
				"/test",
				bytes.NewReader([]byte("a random string")),
				[]*Options{},
			},
			[]byte("a random string"),
			false,
		},
		{
			"should properly issue request with body",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					bdy, _ := io.ReadAll(r.Body)
					w.Write([]byte(bdy))
				},
				"POST",
				"/test",
				bytes.NewBuffer(testData),
				[]*Options{},
			},
			testData,
			false,
		},
	}
	for _, tt := range tests {
		// reset retry count
		retryCount = 0

		// reset the test server
		ts = getTestServer(tt.args.res)
		defer ts.Close()

		u, _ := url.Parse(ts.URL)
		t.Run(tt.name, func(t *testing.T) {
			e := NewEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				"/test",
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := retryRequest(context.Background(), e.a, tt.args.method, tt.args.path, tt.args.body, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			g := strings.TrimSpace(string(got))
			w := strings.TrimSpace(string(tt.want))

			if g != w {
				t.Errorf("Endpoint.request() = \n\t%s\n want \n\t%s", got, tt.want)
			}
		})
	}
}

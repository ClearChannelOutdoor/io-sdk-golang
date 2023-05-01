package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const (
	defaultTestBearerToken string = "test-bearer-token"
	defaultTestEnvironment string = "test"
	defaultTestHost        string = "localhost"
	defaultTestProto       string = "http"
)

func TestEndpoint_Request(t *testing.T) {
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
		opts   []Options
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
				[]Options{},
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
				[]Options{},
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
				[]Options{},
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
				[]Options{},
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
				[]Options{},
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
				[]Options{},
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
			e := NewEndpoint(
				defaultTestBearerToken,
				defaultTestEnvironment,
				u.Host,
				u.Scheme,
			)

			got, err := e.Request(tt.args.method, tt.args.path, tt.args.body, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			g := strings.TrimSpace(string(got))
			w := strings.TrimSpace(string(tt.want))

			if g != w {
				t.Errorf("Endpoint.Request() = \n\t%s\n want \n\t%s", got, tt.want)
			}
		})
	}
}

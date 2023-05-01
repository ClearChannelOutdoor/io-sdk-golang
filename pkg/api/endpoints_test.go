package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
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
		return httptest.NewServer(http.HandlerFunc(res))
	}

	type args struct {
		res    func(w http.ResponseWriter, r *http.Request)
		method string
		path   string
		opts   []Options
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]any
		wantErr bool
	}{
		{
			"should properly succeed in issuing a request",
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")
					io.WriteString(w, `{"test": "test"}`)
				},
				"GET",
				"/test",
				[]Options{},
			},
			map[string]any{"test": "test"},
			false,
		},
	}
	for _, tt := range tests {
		// spin up a test server
		ts := getTestServer(tt.args.res)
		defer ts.Close()

		u, _ := url.Parse(ts.URL)
		t.Run(tt.name, func(t *testing.T) {
			e := NewEndpoint(
				defaultTestBearerToken,
				defaultTestEnvironment,
				u.Host,
				u.Scheme,
			)

			got, err := e.Request(tt.args.method, tt.args.path, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Request() = %v, want %v", got, tt.want)
			}
		})
	}
}

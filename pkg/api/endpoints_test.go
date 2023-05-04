package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	defaultTestBearerToken string = "v2.local.really-long-random-opaque-string"
	defaultTestEnvironment string = "test"
	defaultTestHost        string = "localhost"
	defaultTestProto       string = "http"
)

type TestModel struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

func TestEndpoint_ensureBearerToken(t *testing.T) {
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
					Name:   defaultTestEnvironment,
					Proto:  u.Scheme,
				},
				"/test",
			)

			// set a token url pointing back to our test server
			e.svc.oauth2.TokenURL = fmt.Sprintf("%s/v2/token", ts.URL)
			e.OAuthToken = tt.fields.token

			got, err := e.ensureBearerToken()
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

func TestEndpoint_request(t *testing.T) {
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
			e := NewEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				"/test",
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.request(tt.args.method, tt.args.path, tt.args.body, tt.args.opts...)
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

func TestEndpoint_Create(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
		mdl TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TestModel
		wantErr bool
	}{
		{
			"should properly create a model",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					var tm TestModel
					bdy, _ := io.ReadAll(r.Body)
					json.Unmarshal(bdy, &tm)
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				TestModel{
					ID: "test-model-id",
				},
			},
			TestModel{
				ID:   "test-model-id",
				Path: "/models",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				TestModel{
					ID: "test-model-id",
				},
			},
			TestModel{},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Create(tt.args.mdl)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Create() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_Delete(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should properly delete a model",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNoContent)
				},
				"test-model-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
			},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndpoint_GetOne(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TestModel
		wantErr bool
	}{
		{
			"should properly get a model",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[8:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
			},
			TestModel{
				ID:   "test-model-id",
				Path: "/models/test-model-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
			},
			TestModel{},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.GetOne(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.GetOne() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_GetAll(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SearchResult[TestModel]
		wantErr bool
	}{
		{
			"should properly get all models",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					sr := SearchResult[TestModel]{
						Data: []*TestModel{
							{
								ID:   "test-model-id",
								Path: r.URL.String(),
							},
						},
					}
					json.NewEncoder(w).Encode(sr)
				},
			},
			SearchResult[TestModel]{
				Data: []*TestModel{
					{
						ID:   "test-model-id",
						Path: "/models",
					},
				},
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
			},
			SearchResult[TestModel]{},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.GetAll() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_Patch(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
		id  string
		m   TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TestModel
		wantErr bool
	}{
		{
			"should properly patch a model",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[8:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
				TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			TestModel{
				ID:   "test-model-id",
				Path: "/models/test-model-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			TestModel{},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Patch(tt.args.id, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Patch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Patch() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_Post(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		path string
	}
	type args struct {
		res func(w http.ResponseWriter, r *http.Request)
		id  string
		m   TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    TestModel
		wantErr bool
	}{
		{
			"should properly post a model",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusCreated)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[8:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
				TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			TestModel{
				ID:   "test-model-id",
				Path: "/models/test-model-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := apiError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			TestModel{},
			true,
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
					Host:  u.Host,
					Name:  defaultTestEnvironment,
					Proto: u.Scheme,
				},
				tt.fields.path,
			)

			// set a testing oauth token
			e.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Post(tt.args.id, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Post() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
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
		mdl *TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TestModel
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

					// read the inbound data
					var tm TestModel
					bdy, _ := io.ReadAll(r.Body)
					json.Unmarshal(bdy, &tm)

					// augment with the path
					tm.Path = r.URL.String()
					jsn, _ := json.Marshal(tm)

					// write
					io.WriteString(w, string(jsn))
				},
				&TestModel{
					ID: "test-model-id",
				},
			},
			&TestModel{
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
				&TestModel{
					ID: "test-model-id",
				},
			},
			&TestModel{},
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Create(tt.args.mdl); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.mdl, tt.want) {
				t.Errorf("Endpoint.Create() = \n\t%+v\n want \n\t%+v", tt.args.mdl, tt.want)
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndpoint_Get(t *testing.T) {
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
		want    *TestModel
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
			&TestModel{
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
			&TestModel{},
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Get() = \n\t%+v\n want \n\t%+v", got, tt.want)
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
		m   *TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TestModel
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
				&TestModel{
					ID: "test-model-id",
				},
			},
			&TestModel{
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
				&TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			&TestModel{},
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Patch(tt.args.id, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Patch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("Endpoint.Patch() = \n\t%+v\n want \n\t%+v", tt.args.m, tt.want)
			}
		})
	}
}

func TestEndpoint_Search(t *testing.T) {
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Search()
			if (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Endpoint.Search() = \n\t%+v\n want \n\t%+v", got, tt.want)
			}
		})
	}
}

func TestEndpoint_Update(t *testing.T) {
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
		m   *TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TestModel
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
				&TestModel{
					ID: "test-model-id",
				},
			},
			&TestModel{
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
				&TestModel{
					ID:   "test-model-id",
					Path: "/models/test-model-id",
				},
			},
			&TestModel{},
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
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Update(tt.args.id, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("Endpoint.Update() = \n\t%+v\n want \n\t%+v", tt.args.m, tt.want)
			}
		})
	}
}

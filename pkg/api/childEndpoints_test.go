package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"golang.org/x/oauth2"
)

func TestChildEndpoint_Create(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
		mdl      *TestModel
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TestModel
		wantErr bool
	}{
		{
			"should properly create a child model",
			fields{
				"/models",
				"/children",
				nil,
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
				"test-model-id",
				&TestModel{
					ID: "test-child-id",
				},
			},
			&TestModel{
				ID:   "test-child-id",
				Path: "/models/test-model-id/children",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				&TestModel{
					ID: "test-child-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Create(context.Background(), tt.args.parentID, tt.args.mdl); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.mdl, tt.want) {
				t.Errorf("Endpoint.Create() = \n\t%+v\n want \n\t%+v", tt.args.mdl, tt.want)
			}
		})
	}
}

func TestChildEndpoint_Delete(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
		childID  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should properly delete a child model",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNoContent)
				},
				"test-model-id",
				"test-child-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				"test-child-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Delete(context.Background(), tt.args.parentID, tt.args.childID); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestChildEndpoint_Get(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
		childID  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *TestModel
		wantErr bool
	}{
		{
			"should properly get a child model",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[31:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
				"test-child-id",
			},
			&TestModel{
				ID:   "test-child-id",
				Path: "/models/test-model-id/children/test-child-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				"test-child-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Get(context.Background(), tt.args.parentID, tt.args.childID)
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

func TestChildEndpoint_Patch(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
		id       string
		m        *TestModel
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
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[31:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
				"test-child-id",
				&TestModel{
					ID: "test-child-id",
				},
			},
			&TestModel{
				ID:   "test-child-id",
				Path: "/models/test-model-id/children/test-child-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				"test-child-id",
				&TestModel{
					ID:   "test-child-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Patch(context.Background(), tt.args.parentID, tt.args.id, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Patch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("Endpoint.Patch() = \n\t%+v\n want \n\t%+v", tt.args.m, tt.want)
			}
		})
	}
}

func TestChildEndpoint_Search(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
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
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
					w.Header().Set("Content-Type", "application/json")

					sr := SearchResult[TestModel]{
						Data: []*TestModel{
							{
								ID:   "test-child-id",
								Path: r.URL.String(),
							},
						},
					}
					json.NewEncoder(w).Encode(sr)
				},
				"test-model-id",
			},
			SearchResult[TestModel]{
				Data: []*TestModel{
					{
						ID:   "test-child-id",
						Path: "/models/test-model-id/children",
					},
				},
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			got, err := e.Search(context.Background(), tt.args.parentID)
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

func TestChildEndpoint_Update(t *testing.T) {
	getTestServer := func(res func(w http.ResponseWriter, r *http.Request)) *httptest.Server {
		ts := httptest.NewServer(http.HandlerFunc(res))
		return ts
	}

	var ts *httptest.Server

	type fields struct {
		parentPath string
		childPath  string
		headers    *http.Header
	}
	type args struct {
		res      func(w http.ResponseWriter, r *http.Request)
		parentID string
		id       string
		m        *TestModel
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
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusCreated)
					w.Header().Set("Content-Type", "application/json")

					tm := TestModel{}
					tm.ID = r.URL.Path[31:]
					tm.Path = r.URL.String()
					json.NewEncoder(w).Encode(tm)
				},
				"test-model-id",
				"test-child-id",
				&TestModel{
					ID: "test-child-id",
				},
			},
			&TestModel{
				ID:   "test-child-id",
				Path: "/models/test-model-id/children/test-child-id",
			},
			false,
		},
		{
			"should properly handle an error response",
			fields{
				"/models",
				"/children",
				nil,
			},
			args{
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadRequest)
					w.Header().Set("Content-Type", "application/json")

					err := APIError{
						Message: "test-error-message",
						Status:  http.StatusBadRequest,
						Title:   "test-error-title",
					}

					json.NewEncoder(w).Encode(err)
				},
				"test-model-id",
				"test-child-id",
				&TestModel{
					ID:   "test-child-id",
					Path: "/models/test-model-id/children/test-child-id",
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
			e := NewChildEndpoint[TestModel](
				&Service{
					Host:  u.Host,
					Proto: u.Scheme,
				},
				tt.fields.parentPath,
				tt.fields.childPath,
				tt.fields.headers,
			)

			// set a testing oauth token
			e.a.OAuthToken = &oauth2.Token{
				AccessToken: "test-access-token",
			}

			if err := e.Update(context.Background(), tt.args.parentID, tt.args.id, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("Endpoint.Update() = \n\t%+v\n want \n\t%+v", tt.args.m, tt.want)
			}
		})
	}
}

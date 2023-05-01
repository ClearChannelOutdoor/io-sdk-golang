// The api package contains the API drivers for the each of the
// various IO platform APIs (microservices) that are supported.
// Use of this package is intended to ease development and ensure
// best practices are followed with regards to token management,
// error handling, and other common API concerns including retry logic,
// and pagination.
package api

import "testing"

func Test_containsValue(t *testing.T) {
	type args struct {
		values []string
		value  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{[]string{}, ""}, false},
		{"not found", args{[]string{"a", "b", "c"}, "d"}, false},
		{"found", args{[]string{"a", "b", "c"}, "b"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsValue(tt.args.values, tt.args.value); got != tt.want {
				t.Errorf("containsValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

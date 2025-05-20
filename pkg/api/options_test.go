package api

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_options_Options(t *testing.T) {
	o := EmptyOptions()

	if o == nil {
		t.Errorf("options.Options() = nil")
	}
}

func Test_options_AddField(t *testing.T) {
	type fields struct {
		fields []string
	}
	type args struct {
		field []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should add a field",
			fields{
				fields: []string{},
			},
			args{
				[]string{"field"},
			},
			&Options{
				fields: []string{"field"},
			},
		},
		{
			"should not double add a field",
			fields{
				fields: []string{"field"},
			},
			args{
				[]string{"field"},
			},
			&Options{
				fields: []string{"field"},
			},
		},
		{
			"should add multiple fields",
			fields{
				fields: []string{"field 1"},
			},
			args{
				[]string{"field 2", "field 3"},
			},
			&Options{
				fields: []string{"field 1", "field 2", "field 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				fields: tt.fields.fields,
			}
			if got := o.AddField(tt.args.field...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.AddField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_RemoveField(t *testing.T) {
	type fields struct {
		fields []string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should remove a field",
			fields{
				fields: []string{"field"},
			},
			args{
				"field",
			},
			&Options{
				fields: []string{},
			},
		},
		{
			"should not remove a field not present",
			fields{
				fields: []string{"field 2"},
			},
			args{
				"field 1",
			},
			&Options{
				fields: []string{"field 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				fields: tt.fields.fields,
			}
			if got := o.RemoveField(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.RemoveField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_AddFilter(t *testing.T) {
	type fields struct {
		filter map[string][]string
	}
	type args struct {
		field string
		value []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should add a filter with a string",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{"value"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value"},
				},
			},
		},
		{
			"should add a filter with an integer",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{10},
			},
			&Options{
				filter: map[string][]string{
					"field": {"10"},
				},
			},
		},
		{
			"should add a filter with a fmt.Stringer (time.Time)",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{time.Date(1, 1, 1, 12, 0, 0, 0, time.UTC)},
			},
			&Options{
				filter: map[string][]string{
					"field": {"0001-01-01 12:00:00 +0000 UTC"},
				},
			},
		},
		{
			"should add additional values to existing filter",
			fields{
				filter: map[string][]string{
					"field": {"value 1"},
				},
			},
			args{
				"field",
				[]any{"value 2"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
		},
		{
			"should add multiple values to filter when passed as array",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{"value 1", "value 2"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
		},
		{
			"should not duplicate values that already exist",
			fields{
				filter: map[string][]string{
					"field": {"value 1"},
				},
			},
			args{
				"field",
				[]any{"value 1"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value 1"},
				},
			},
		},
		{
			"should add multiple filter values when provided variadically",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{"value-1", "value-2", "value-3"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value-1", "value-2", "value-3"},
				},
			},
		},
		{
			"should add multiple filter values when provided variadically (even with duplicate)",
			fields{
				filter: map[string][]string{},
			},
			args{
				"field",
				[]any{"value-1", "value-2", "value-2", "value-3"},
			},
			&Options{
				filter: map[string][]string{
					"field": {"value-1", "value-2", "value-3"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Options{
				filter: tt.fields.filter,
			}

			if got := o.AddFilter(tt.args.field, tt.args.value...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.AddFilter() = \n\t%+v\nwant \n\t%+v", got, tt.want)
			}
		})
	}
}

func Test_options_RemoveFilter(t *testing.T) {
	type fields struct {
		filter map[string][]string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should remove a filter",
			fields{
				filter: map[string][]string{
					"field": {"value"},
				},
			},
			args{
				"field",
			},
			&Options{
				filter: map[string][]string{},
			},
		},
		{
			"should remove a filter with multiple values",
			fields{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
			args{
				"field",
			},
			&Options{
				filter: map[string][]string{},
			},
		},
		{
			"should not remove an unmatched filter field",
			fields{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
			args{
				"nomatch",
			},
			&Options{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				filter: tt.fields.filter,
			}
			if got := o.RemoveFilter(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.RemoveFilter() = \n\t%+v\nwant \n\t%+v", got, tt.want)
			}
		})
	}
}

func Test_options_AddPage(t *testing.T) {
	type fields struct {
		page map[string]int
	}
	type args struct {
		field string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should add a field",
			fields{
				page: map[string]int{},
			},
			args{
				"limit",
				1000,
			},
			&Options{
				page: map[string]int{
					"limit": 1000,
				},
			},
		},
		{
			"should not double add a field",
			fields{
				page: map[string]int{
					"offset": 0,
				},
			},
			args{
				"offset",
				0,
			},
			&Options{
				page: map[string]int{
					"offset": 0,
				},
			},
		},
		{
			"should properly update a field",
			fields{
				page: map[string]int{
					"limit": 100,
				},
			},
			args{
				"limit",
				500,
			},
			&Options{
				page: map[string]int{
					"limit": 500,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				page: tt.fields.page,
			}
			if got := o.AddPage(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.AddPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_AddSort(t *testing.T) {
	type fields struct {
		sort []string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should add a field",
			fields{
				sort: []string{},
			},
			args{
				"field",
			},
			&Options{
				sort: []string{"field"},
			},
		},
		{
			"should not double add a field",
			fields{
				sort: []string{"field"},
			},
			args{
				"field",
			},
			&Options{
				sort: []string{"field"},
			},
		},
		{
			"should add a field to the end",
			fields{
				sort: []string{"field 1"},
			},
			args{
				"field 2",
			},
			&Options{
				sort: []string{"field 1", "field 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				sort: tt.fields.sort,
			}
			if got := o.AddSort(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.AddSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_NextPage(t *testing.T) {
	type fields struct {
		page map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Options
	}{
		{
			"should increment the offset by the limit",
			fields{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 100,
				},
			},
		},
		{
			"should increment the offset by the default limit when limit is not set",
			fields{
				page: map[string]int{
					"offset": 0,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 100,
				},
			},
		},
		{
			"should not increment the offset when limit is 0",
			fields{
				page: map[string]int{
					"limit":  0,
					"offset": 0,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  0,
					"offset": 0,
				},
			},
		},
		{
			"should not increment the offset when limit is less than 0",
			fields{
				page: map[string]int{
					"limit":  -1,
					"offset": 0,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  -1,
					"offset": 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				page: tt.fields.page,
			}
			if got := o.NextPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.NextPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_PageDefault(t *testing.T) {
	type fields struct {
		page map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Options
	}{
		{
			"should return default pagination options",
			fields{
				page: map[string]int{},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
		},
		{
			"should clear extraneous pagination options",
			fields{
				page: map[string]int{
					"page": 1,
					"size": 100,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				page: tt.fields.page,
			}
			if got := o.PageDefault(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.DefaultPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_PreviousPage(t *testing.T) {
	type fields struct {
		page map[string]int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Options
	}{
		{
			"should decrement the offset by the limit",
			fields{
				page: map[string]int{
					"limit":  100,
					"offset": 100,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
		},
		{
			"should decrement the offset by the default limit when limit is not set",
			fields{
				page: map[string]int{
					"offset": 100,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
		},
		{
			"should not decrement the offset when limit is 0",
			fields{
				page: map[string]int{
					"limit":  0,
					"offset": 100,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  0,
					"offset": 100,
				},
			},
		},
		{
			"should not decrement the offset when limit is less than 0",
			fields{
				page: map[string]int{
					"limit":  -1,
					"offset": 100,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  -1,
					"offset": 100,
				},
			},
		},
		{
			"should not decrement the offset when offset is less than 0",
			fields{
				page: map[string]int{
					"limit":  100,
					"offset": -1,
				},
			},
			&Options{
				page: map[string]int{
					"limit":  100,
					"offset": 0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				page: tt.fields.page,
			}

			if got := o.PreviousPage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.PreviousPage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_RemovePage(t *testing.T) {
	type fields struct {
		page map[string]int
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should remove a field",
			fields{
				page: map[string]int{
					"limit": 100,
				},
			},
			args{
				"limit",
			},
			&Options{
				page: map[string]int{},
			},
		},
		{
			"should not remove an unmatched field",
			fields{
				page: map[string]int{
					"limit": 100,
				},
			},
			args{
				"offset",
			},
			&Options{
				page: map[string]int{
					"limit": 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				page: tt.fields.page,
			}
			if got := o.RemovePage(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.RemovePage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_RemoveSort(t *testing.T) {
	type fields struct {
		sort []string
	}
	type args struct {
		field string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Options
	}{
		{
			"should remove a field",
			fields{
				sort: []string{"field"},
			},
			args{
				"field",
			},
			&Options{
				sort: []string{},
			},
		},
		{
			"should not remove an unmatched field",
			fields{
				sort: []string{"field"},
			},
			args{
				"other",
			},
			&Options{
				sort: []string{"field"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				sort: tt.fields.sort,
			}
			if got := o.RemoveSort(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.RemoveSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_Query(t *testing.T) {
	type fields struct {
		fields []string
		filter map[string][]string
		page   map[string]int
		sort   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		{
			"should return an empty query when no options are set",
			fields{},
			url.Values{},
		},
		{
			"should return a query with fields provided",
			fields{
				fields: []string{"field"},
			},
			url.Values{
				"fields": []string{"field"},
			},
		},
		{
			"should return a query with filters provided",
			fields{
				filter: map[string][]string{
					"field": {"value"},
					"other": {"value"},
				},
			},
			url.Values{
				"filter[field]": []string{"value"},
				"filter[other]": []string{"value"},
			},
		},
		{
			"should return a query with filters comma separated when there are multiple values",
			fields{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
			url.Values{
				"filter[field]": []string{"value 1", "value 2"},
			},
		},
		{
			"should return a query with page provided",
			fields{
				page: map[string]int{
					"limit":  100,
					"offset": 100,
				},
			},
			url.Values{
				"page[limit]":  []string{"100"},
				"page[offset]": []string{"100"},
			},
		},
		{
			"should return a query with sort provided",
			fields{
				sort: []string{"field", "-other"},
			},
			url.Values{
				"sort": []string{"field", "-other"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				fields: tt.fields.fields,
				filter: tt.fields.filter,
				page:   tt.fields.page,
				sort:   tt.fields.sort,
			}

			if got := o.Query(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("options.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_FormatQuery(t *testing.T) {
	var similar = func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}

		aPrms := strings.Split(a, "&")
		bPrms := strings.Split(b, "&")

		if len(aPrms) != len(bPrms) {
			return false
		}

		for _, prm := range aPrms {
			if !strings.Contains(b, prm) {
				return false
			}
		}

		return true
	}
	type fields struct {
		fields []string
		filter map[string][]string
		page   map[string]int
		sort   []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"should return an empty query when no options are set",
			fields{},
			"",
		},
		{
			"should return a query with fields provided",
			fields{
				fields: []string{"field"},
			},
			"fields=field",
		},
		{
			"should return a query with filters provided",
			fields{
				filter: map[string][]string{
					"field": {"value"},
					"other": {"value"},
				},
			},
			"filter[field]=value&filter[other]=value",
		},
		{
			"should return a query with filters comma separated when there are multiple values",
			fields{
				filter: map[string][]string{
					"field": {"value 1", "value 2"},
				},
			},
			"filter[field]=value+1,value+2",
		},
		{
			"should return a query with page provided",
			fields{
				page: map[string]int{
					"limit":  100,
					"offset": 100,
				},
			},
			"page[limit]=100&page[offset]=100",
		},
		{
			"should return a query with sort provided",
			fields{
				sort: []string{"field", "-other"},
			},
			"sort=field,-other",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Options{
				fields: tt.fields.fields,
				filter: tt.fields.filter,
				page:   tt.fields.page,
				sort:   tt.fields.sort,
			}

			if got := o.FormatQuery(); !similar(got, tt.want) {
				t.Errorf("options.FormatQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

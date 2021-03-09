package httpassert

import (
	"reflect"
	"testing"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type paginationData struct {
	Total    uint64      `json:"total"`
	Offset   int64       `json:"offset"`
	Limit    int64       `json:"limit"`
	Previous string      `json:"previous"`
	Next     string      `json:"next"`
	Items    interface{} `json:"items"`
}

type user struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
}

func TestCompareValues(t *testing.T) {
	r1 := response{
		Code:    0,
		Message: "ok",
		Data: &paginationData{
			Total:    1234,
			Offset:   30,
			Limit:    15,
			Previous: "http://localhost/previous",
			Next:     "http://localhost/next",
			Items: []*user{
				{
					Name:   "John",
					Age:    18,
					Height: 182.3,
				},
				{
					Name:   "Lily",
					Age:    22,
					Height: 170.8,
				},
			},
		},
	}

	r2 := response{
		Code:    0,
		Message: "ok",
		Data: map[string]interface{}{
			"total":    1234,
			"offset":   30,
			"limit":    15,
			"previous": "http://localhost/previous",
			"next":     "http://localhost/next",
			"items": []map[string]interface{}{
				{
					"name":   "John",
					"age":    18,
					"height": 182.3,
				},
				{
					"name":   "Lily",
					"age":    22,
					"height": 170.8,
				},
			},
		},
	}
	tests := []struct {
		arg1     interface{}
		arg2     interface{}
		expected bool
	}{
		{
			arg1:     r1,
			arg2:     r2,
			expected: true,
		},
		{
			arg1:     map[string]string{"ni": "hao", "hello": "world"},
			arg2:     map[string]string{"hello": "world", "ni": "hao"},
			expected: true,
		},
		{
			arg1:     complex(1.0, 1.0),
			arg2:     complex(1.0, 1.0),
			expected: true,
		},
		{
			arg1:     complex64(complex(1.0, 1.0)),
			arg2:     complex64(complex(1.0, 1.0)),
			expected: true,
		},
		{
			arg1:     complex(1.0, 1.0),
			arg2:     complex64(complex(1.0, 1.0)),
			expected: false,
		},
		{
			// not support compare functions
			arg1:     func() {},
			arg2:     func() {},
			expected: false,
		},
		{
			arg1: paginationData{
				Total:  1,
				Offset: 2,
				Limit:  10,
				Items:  Any{},
			},
			arg2: paginationData{
				Total:  1,
				Offset: 2,
				Limit:  10,
				Items:  "hello world",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		actual := compareValues(reflect.ValueOf(tt.arg1), reflect.ValueOf(tt.arg2))

		if actual != tt.expected {
			t.Errorf("compareValues((%T) %v, (%T) %v) got unexpected result: want %v, got %v", tt.arg1, tt.arg1, tt.arg2, tt.arg2, tt.expected, actual)
		}
	}
}

package httpassert

import (
	"reflect"
	"testing"
	"time"
)

func TestRenderValue(t *testing.T) {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	var nilV interface{}

	tests := []struct {
		val      interface{}
		expected string
	}{
		{
			val:      nilV,
			expected: "nil",
		},
		{
			val:      1,
			expected: "1",
		},
		{
			val:      int8(1),
			expected: "1",
		},
		{
			val:      int16(1),
			expected: "1",
		},
		{
			val:      int32(1),
			expected: "1",
		},
		{
			val:      int64(1),
			expected: "1",
		},
		{
			val:      uint(1),
			expected: "1",
		},
		{
			val:      uint8(1),
			expected: "1",
		},
		{
			val:      uint16(1),
			expected: "1",
		},
		{
			val:      uint32(1),
			expected: "1",
		},
		{
			val:      uint64(1),
			expected: "1",
		},
		{
			val:      true,
			expected: "true",
		},
		{
			val:      false,
			expected: "false",
		},
		{
			val:      1.234,
			expected: "1.234",
		},
		{
			val:      float32(1.234),
			expected: "1.234",
		},
		{
			val:      complex(1.0, 1.0),
			expected: "(1+1i)",
		},
		{
			val:      complex64(complex(3.14, -8.0)),
			expected: "(3.14-8i)",
		},
		{
			val:      "ok",
			expected: `"ok"`,
		},
		{
			val:      []int{1, 2, 3},
			expected: "[1, 2, 3]",
		},
		{
			val:      []string{"foo", "bar", "fuz"},
			expected: `["foo", "bar", "fuz"]`,
		},
		{
			val: []user{
				{
					Name:   "John",
					Age:    20,
					Height: 178,
				},
				{
					Name:   "Lily",
					Age:    20,
					Height: 170,
				},
			},
			expected: `[
    {
        Name: (string) "John",
        Age: (int) 20,
        Height: (float64) 178,
        Extra: (interface{}) nil,
    },
    {
        Name: (string) "Lily",
        Age: (int) 20,
        Height: (float64) 170,
        Extra: (interface{}) nil,
    },
]`,
		},
		{
			val: [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: `[
    [1, 2, 3],
    [4, 5, 6],
]`,
		},
		{
			val: []map[interface{}]interface{}{
				{
					"name": "John",
				},
				{
					"name": "Lily",
				},
			},
			expected: `[
    {
        "name": (string) "John",
    },
    {
        "name": (string) "Lily",
    },
]`,
		},
		{
			val:      time.Date(2020, 1, 1, 10, 0, 0, 0, beijing),
			expected: "2020-01-01 10:00:00 +0800 Beijing Time",
		},
	}

	for _, tt := range tests {
		actual := renderValue(reflect.ValueOf(tt.val), 0)

		if actual != tt.expected {
			t.Errorf(
				"renderValue((%T) %v) unexpected result: want %s, got %s",
				tt.val, tt.val, tt.expected, actual,
			)
		}
	}
}

package httpassert

import (
	"reflect"
	"testing"
)

func TestRenderValue(t *testing.T) {
	tests := []struct {
		val      interface{}
		expected string
	}{
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
    },
    {
        Name: (string) "Lily",
        Age: (int) 20,
        Height: (float64) 170,
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
		/* skip test map, since the keys is not sorted
				{
					val: []map[string]interface{}{
						{
							"name":   "John",
							"age":    20,
							"height": 178,
						},
						{
							"name":   "Lily",
							"age":    20,
							"height": 170,
						},
					},
					expected: `[
		    {
		        "name": "John",
		        "age": 20,
		        "height": 178,
		    },
		    {
		        "name": "Lily",
		        "age": 20,
		        "height": 170,
		    },
		]`,
				},
		*/
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
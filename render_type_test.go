package httpassert

import (
	"reflect"
	"testing"
	"time"
)

func TestRenderType(t *testing.T) {
	var nilV interface{}

	tests := []struct {
		typ      reflect.Type
		expected string
	}{
		{
			typ:      reflect.TypeOf(nilV),
			expected: "interface{}",
		},
		{
			typ:      reflect.TypeOf(1),
			expected: "int",
		},
		{
			typ:      reflect.TypeOf(int8(1)),
			expected: "int8",
		},
		{
			typ:      reflect.TypeOf(int16(1)),
			expected: "int16",
		},
		{
			typ:      reflect.TypeOf(int32(1)),
			expected: "int32",
		},
		{
			typ:      reflect.TypeOf(int64(1)),
			expected: "int64",
		},
		{
			typ:      reflect.TypeOf(uint(1)),
			expected: "uint",
		},
		{
			typ:      reflect.TypeOf(uint8(1)),
			expected: "uint8",
		},
		{
			typ:      reflect.TypeOf(uint16(1)),
			expected: "uint16",
		},
		{
			typ:      reflect.TypeOf(uint32(1)),
			expected: "uint32",
		},
		{
			typ:      reflect.TypeOf(uint64(1)),
			expected: "uint64",
		},
		{
			typ:      reflect.TypeOf(true),
			expected: "bool",
		},
		{
			typ:      reflect.TypeOf("hello"),
			expected: "string",
		},
		{
			typ:      reflect.TypeOf(1.0),
			expected: "float64",
		},
		{
			typ:      reflect.TypeOf(float32(1.0)),
			expected: "float32",
		},
		{
			typ:      reflect.TypeOf(complex(1.0, 1.0)),
			expected: "complex128",
		},
		{
			typ:      reflect.TypeOf(complex64(complex(1.0, 1.0))),
			expected: "complex64",
		},
		{
			typ:      reflect.TypeOf(&user{}),
			expected: "*github.com/zgs225/httpassert.user",
		},
		{
			typ:      reflect.TypeOf(user{}),
			expected: "github.com/zgs225/httpassert.user",
		},
		{
			typ:      reflect.TypeOf(struct{}{}),
			expected: "anonymous",
		},
		{
			typ:      reflect.TypeOf(map[string]interface{}{}),
			expected: "map[string]interface",
		},
		{
			typ:      reflect.TypeOf(map[interface{}]interface{}{}),
			expected: "map[interface]interface",
		},
		{
			typ:      reflect.TypeOf(map[interface{}]*user{}),
			expected: "map[interface]*github.com/zgs225/httpassert.user",
		},
		{
			typ:      reflect.TypeOf(make(chan *user)),
			expected: "chan *github.com/zgs225/httpassert.user",
		},
		{
			typ:      reflect.TypeOf(time.Now()),
			expected: "time.Time",
		},
	}

	for _, tt := range tests {
		if actual := renderType(tt.typ); actual != tt.expected {
			t.Errorf(
				"renderType(%v) unexpected result: want %s, got %s",
				tt.typ, tt.expected, actual,
			)
		}
	}
}

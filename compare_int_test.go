package httpassert

import (
	"reflect"
	"testing"
)

func TestCompareInt(t *testing.T) {
	tests := []testCase{
		{
			v1:       1,
			v2:       1,
			expected: true,
		},
		{
			v1:       1,
			v2:       -1,
			expected: false,
		},
		{
			v1:       int8(1),
			v2:       int8(1),
			expected: true,
		},
		{
			v1:       int16(1),
			v2:       int16(1),
			expected: true,
		},
		{
			v1:       int32(1),
			v2:       int32(1),
			expected: true,
		},
		{
			v1:       int64(1),
			v2:       int64(1),
			expected: true,
		},
	}

	runTestCases(t, tests, compareInt)
}

func runTestCases(t *testing.T, cases []testCase, fn interface{}) {
	fnv := reflect.ValueOf(fn)

	for _, tt := range cases {
		args := []reflect.Value{
			valueOfTwice(tt.v1),
			valueOfTwice(tt.v2),
		}

		if fnv.Type().NumIn() == 3 {
			args = append(args, reflect.ValueOf(reflect.ValueOf(tt.v1).Kind()))
		}

		ret := fnv.Call(args)

		actual := ret[0].Bool()

		if actual != tt.expected {
			t.Errorf("compare(%T(%v), %T(%v)) error: want %T(%v), got %T(%v)\n", tt.v1, tt.v1, tt.v2, tt.v2, tt.expected, tt.expected, actual, actual)
		}
	}
}

func valueOfTwice(i interface{}) reflect.Value {
	return reflect.ValueOf(reflect.ValueOf(i))
}

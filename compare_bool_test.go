package httpassert

import (
	"testing"
)

type testCase struct {
	v1       interface{}
	v2       interface{}
	expected bool
}

func TestCompareBool(t *testing.T) {
	tests := []testCase{
		{
			v1:       true,
			v2:       true,
			expected: true,
		},
		{
			v1:       false,
			v2:       false,
			expected: true,
		},
		{
			v1:       true,
			v2:       false,
			expected: false,
		},
		{
			v1:       false,
			v2:       true,
			expected: false,
		},
		{
			v1:       true,
			v2:       "3",
			expected: false,
		},
	}

	runTestCases(t, tests, compareBool)
}

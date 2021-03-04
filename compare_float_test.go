package httpassert

import "testing"

func TestCompareFloat(t *testing.T) {
	tests := []testCase{
		{
			v1:       1.0,
			v2:       1.0,
			expected: true,
		},
		{
			v1:       float32(2.0),
			v2:       float32(2.0),
			expected: true,
		},
		{
			v1:       3.3,
			v2:       2.2,
			expected: false,
		},
		{
			v1:       float64(1.0),
			v2:       float32(1.0),
			expected: false,
		},
	}

	runTestCases(t, tests, compareFloat)
}

package httpassert

import "testing"

func TestCompareString(t *testing.T) {
	tests := []testCase{
		{
			v1:       "ok",
			v2:       "ok",
			expected: true,
		},
		{
			v1:       "ok",
			v2:       "not ok",
			expected: false,
		},
		{
			v1:       "1",
			v2:       1,
			expected: false,
		},
	}

	runTestCases(t, tests, compareString)
}

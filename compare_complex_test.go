package httpassert

import "testing"

func TestCompareComplex(t *testing.T) {
	tests := []testCase{
		{
			v1:       complex(1.0, -2.0),
			v2:       complex(1.0, -2.0),
			expected: true,
		},
		{
			v1:       complex(1.0, -2.0),
			v2:       complex(-2.0, 1.0),
			expected: false,
		},
	}

	runTestCases(t, tests, compareComplex)
}

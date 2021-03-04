package httpassert

import "testing"

func TestCompareArrayLike(t *testing.T) {
	tests := []testCase{
		{
			v1:       []int{1, 2, 3},
			v2:       []int{1, 2, 3},
			expected: true,
		},
		{
			v1:       []int{1, 2, 3},
			v2:       []int{3, 2, 1},
			expected: false,
		},
		{
			v1:       []int{1, 2, 3},
			v2:       []int64{1, 2, 3},
			expected: false,
		},
		{
			v1:       []interface{}{1, "ok", true, 3.0},
			v2:       []interface{}{1, "ok", true, 3.0},
			expected: true,
		},
		{
			v1:       [3]int{1, 2, 3},
			v2:       []int{1, 2, 3},
			expected: false,
		},
		{
			v1:       [3]int{1, 2, 3},
			v2:       [3]int{1, 2, 3},
			expected: true,
		},
	}

	runTestCases(t, tests, compareArrayLike)
}

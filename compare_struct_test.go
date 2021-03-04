package httpassert

import "testing"

func TestCompareStruct(t *testing.T) {
	tests := []testCase{
		{
			v1:       user{},
			v2:       user{},
			expected: true,
		},
		{
			v1: user{
				Name:   "lily",
				Age:    18,
				Height: 170.0,
			},
			v2: user{
				Name:   "lily",
				Age:    18,
				Height: 170.0,
			},
			expected: true,
		},
		{
			v1: user{
				Name:   "lily",
				Age:    18,
				Height: 170.0,
			},
			v2: map[string]interface{}{
				"name":   "lily",
				"age":    18,
				"height": 170.0,
			},
			expected: true,
		},
		{
			v1: user{
				Name:   "john",
				Age:    18,
				Height: 170.0,
			},
			v2: user{
				Name:   "lily",
				Age:    18,
				Height: 170.0,
			},
			expected: false,
		},
	}

	runTestCases(t, tests, compareStruct)
}

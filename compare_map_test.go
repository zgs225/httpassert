package httpassert

import "testing"

func TestCompareMap(t *testing.T) {
	tests := []testCase{
		{
			v1:       map[string]string(nil),
			v2:       map[string]string(nil),
			expected: true,
		},
		{
			v1: map[string]string{
				"hello": "world",
				"ni":    "hao",
			},
			v2: map[string]interface{}{
				"hello": "world",
				"ni":    "hao",
			},
			expected: true,
		},
		{
			v1: map[string]interface{}{
				"name":   "Zhang san",
				"age":    20,
				"height": 178.3,
			},
			v2: user{
				Name:   "Zhang san",
				Age:    20,
				Height: 178.3,
			},
			expected: true,
		},
		{
			v1:       map[string]string(nil),
			v2:       "hello",
			expected: false,
		},
	}

	runTestCases(t, tests, compareMap)
}

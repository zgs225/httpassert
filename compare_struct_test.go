package httpassert

import (
	"testing"
	"time"
)

func TestCompareStruct(t *testing.T) {
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

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
		{
			v1:       time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
			v2:       time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
			expected: true,
		},
		{
			v1:       time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
			v2:       time.Date(2000, 2, 1, 20, 30, 0, 0, beijing),
			expected: true,
		},
		{
			v1:       time.Date(2020, 1, 1, 10, 0, 0, 0, time.Local),
			v2:       time.Date(2021, 1, 1, 10, 0, 0, 0, time.Local),
			expected: false,
		},
	}

	runTestCases(t, tests, compareStruct)
}

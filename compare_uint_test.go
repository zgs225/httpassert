package httpassert

import (
	"testing"
)

func TestCompareUint(t *testing.T) {
	tests := []testCase{
		{
			v1:       uint(1),
			v2:       uint(1),
			expected: true,
		},
		{
			v1:       uint(1),
			v2:       uint(0),
			expected: false,
		},
		{
			v1:       uint8(1),
			v2:       uint8(1),
			expected: true,
		},
		{
			v1:       uint16(1),
			v2:       uint16(1),
			expected: true,
		},
		{
			v1:       uint32(1),
			v2:       uint32(1),
			expected: true,
		},
		{
			v1:       uint64(1),
			v2:       uint64(1),
			expected: true,
		},
	}

	runTestCases(t, tests, compareUint)
}

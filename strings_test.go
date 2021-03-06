package httpassert

import (
	"strconv"
	"testing"
)

func TestNindent(t *testing.T) {
	tests := []struct {
		s        string
		n        int
		expected string
	}{
		{
			s:        "hello world",
			n:        1,
			expected: "    hello world",
		},
		{
			s:        "hello\nworld\n",
			n:        2,
			expected: "        hello\n        world\n",
		},
		{
			s:        "hello\nworld",
			n:        2,
			expected: "        hello\n        world",
		},
	}

	for _, tt := range tests {
		actual := nindent(tt.s, indentStr, tt.n)

		if actual != tt.expected {
			t.Errorf(
				"nindent(%s, %s, %d) error: want %s, got %s",
				strconv.Quote(tt.s), strconv.Quote(indentStr), tt.n,
				tt.expected, actual,
			)
		}
	}
}

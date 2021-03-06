package httpassert

import (
	"bytes"
	"strings"
)

const indentStr = "    "

func nindent(s, indent string, n int) string {
	if n <= 0 {
		return s
	}

	buf := new(bytes.Buffer)
	parts := strings.Split(s, "\n")

	for _, part := range parts {
		if len(part) > 0 {
			buf.WriteString(strings.Repeat(indent, n))
			buf.WriteString(part)
		}
		buf.WriteByte('\n')
	}

	buf.Truncate(buf.Len() - 1)

	return buf.String()
}

package comma

import (
	"bytes"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaNotRecursive(s string) string {
	ost := len(s) % 3
	if ost == 0 {
		ost = 3
	}
	buf := bytes.Buffer{}
	buf.WriteString(s[:ost])
	for i := ost; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}
	return buf.String()
}

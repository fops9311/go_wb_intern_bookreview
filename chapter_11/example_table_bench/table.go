package table

import (
	"bytes"
	"strings"
)

func AppendStringsPlus(s1, s2 string) string {
	return s1 + s2
}

func AppendStringsAppend(s1, s2 string) string {
	return string(append([]byte(s1), []byte(s2)...))
}

func AppendStringsBuffer(s1, s2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(s1)
	buffer.WriteString(s2)
	return buffer.String()
}
func AppendStringsBuilder(s1, s2 string) string {
	b := strings.Builder{}
	//b.Grow(len(s1) + len(s2))
	b.WriteString(s1)
	b.WriteString(s2)
	return b.String()
}

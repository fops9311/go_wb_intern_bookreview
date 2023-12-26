package table

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var result string

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func Seq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[0]
	}
	return string(b)
}

var sls = func() []string {
	res := make([]string, 5000000)
	for i := 0; i < len(res); i++ {
		res = append(res, randSeq(10))
	}
	return res
}()

func BenchmarkAppendStringsPlus(b *testing.B) {
	var s string
	for i := 0; i < b.N; i = i + 1 {
		s += sls[i]
	}
	result = s
}
func BenchmarkAppendStringsAppend(b *testing.B) {
	var s string
	for i := 0; i < b.N; i = i + 1 {
		s = AppendStringsAppend(s, sls[i])
	}
	result = s
}
func BenchmarkAppendStringsBuffer(b *testing.B) {
	var s string
	for i := 0; i < b.N; i = i + 1 {
		s = AppendStringsBuffer(s, sls[i])
	}
	result = s
}

func BenchmarkAppendStringsBuilder(b *testing.B) {
	var s string
	for i := 0; i < b.N; i = i + 1 {
		s = AppendStringsBuilder(s, sls[i])
	}
	result = s
}
func BenchmarkAppendStringsFmt(b *testing.B) {
	var s string
	for i := 0; i < b.N; i = i + 1 {
		s = fmt.Sprintf("%s%s", s, sls[i])
	}
	result = s
}

func BenchmarkConstBuilder(b *testing.B) {
	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString(sls[n])
	}
	result = sb.String()
}
func BenchmarkConstBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString(sls[n])
	}
	result = buffer.String()
}

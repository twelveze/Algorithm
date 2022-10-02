package BenchMark

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkBuffer(b *testing.B) {
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&buf, "123")
		_ = buf.String()
	}
}

func BenchmarkBuilder(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < b.N; i++ {
		fmt.Fprint(&builder, "123")
		_ = builder.String()
	}
}

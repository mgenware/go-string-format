package strf

import (
	"fmt"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("🙅‍♂️%v🧛‍♂️%v%v", 1, "haha", -32)
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format("🙅‍♂️{0}🧛‍♂️{1}{2}", 1, "haha", -32)
	}
}

func BenchmarkFormatCore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Format("🙅‍♂️{}🧛‍♂️{}{}", 1, "haha", -32)
	}
}

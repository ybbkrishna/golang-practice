package lockvsparse

import (
	"testing"
)

func Benchmark_parse_always(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse_always()
	}
}

func Benchmark_parse_once_with_lock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parse_once_with_lock()
	}
}

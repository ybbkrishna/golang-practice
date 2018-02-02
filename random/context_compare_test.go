package main

import (
	"context"
	"testing"
)

func BenchmarkFuncByCtx(b *testing.B) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, Item, "h")
	for i := 0; i < b.N; i++ {
		funcByCtx(ctx)
	}
}

func BenchmarkFuncByVal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funcByVal("h")
	}
}

package main

import (
	"context"
)

const (
	Item = iota
)

func funcByCtx(ctx context.Context) {
	val := ctx.Value(Item)
	str := val.(string)
	if str == "xx" {
	}
}

func funcByVal(str string) {
	if str == "xx" {
	}
}

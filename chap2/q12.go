package main

import (
	"fmt"
)

func mapper(f func(interface{}) interface{}, vals []interface{}) []interface{} {
	out := make([]interface{}, len(vals))
	for i, val := range vals {
		out[i] = f(val)
	}
	return out
}

func square(i interface{}) interface{} {
	val, ok := i.(int)
	if ok {
		return val * val
	}
	panic("Not a number")
}

func main() {
	arr := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(mapper(square, arr))
}

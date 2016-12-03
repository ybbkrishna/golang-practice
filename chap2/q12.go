/*
Map function A map()-function is a function that takes a function and a list.
The function is applied to each member in the list and a new list containing these calculated values is returned.
Thus: map(f(), (a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))
1. Write a simple map()-function in Go. It is sufficient for this function only to work
for ints.
2. Expand your code to also work on a list of strings.
*/
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

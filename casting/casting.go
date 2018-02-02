package main

import (
	"fmt"
)

func main() {
	var x interface{} = nil
	val, ok := x.(bool)
	fmt.Print(val, ok)
	x = false
	val, ok = x.(bool)
	fmt.Print(val, ok)
}

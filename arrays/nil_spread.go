package main

import (
	"fmt"
)

func main() {
	var x []int
	x = nil
	lol := []int{1, 2, 3}
	l := append(lol, x...)
	fmt.Println(l)
}

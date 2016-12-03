package main

import (
	"fmt"
)

func plus(i int) func(j int) int {
	return func(j int) int {
		return i + j
	}
}

func main() {
	fmt.Println(plus(2)(3))
}

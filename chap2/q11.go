package main

import (
	"fmt"
)

func fibonacci(count int) {
	a, b := 1, 1
	fmt.Printf("%d ", a)
	fmt.Printf("%d ", b)
	for i := 1; i < count-1; i++ {
		a, b = b, a+b
		fmt.Printf("%d ", b)
	}
}

func main() {
	fibonacci(20)
}

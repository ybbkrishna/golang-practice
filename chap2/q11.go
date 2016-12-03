/*
The Fibonacci sequence starts as follows :1,1,2,3,5,8,13,...Or in mathematical
terms: x1 = 1;x2 = 1;xn = xn−1 +xn−2 ∀n > 2.
Write a function that takes an int value and gives that many terms of the Fibonacci
sequence.
*/

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

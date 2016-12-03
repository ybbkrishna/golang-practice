/*

1. Write a function that returns a function that performs a+2 on integers.
Name the function plusTwo. You should then be able do the following: p := plusTwo()
      fmt.Printf("%v\n", p(2))
Which should print 4. See section Callbacks on page 29 for information about this topic.
2. Generalize the function from 1, and create a plusX(x) which returns functions that add x to an integer.
*/
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

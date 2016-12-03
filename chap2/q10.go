package main

import (
	"fmt"
)

func variadicFunction(args ...int) {
	for _, val := range args {
		fmt.Println(val)
	}
}

func main() {
	variadicFunction(1, 2, 34, 5, 8)
}

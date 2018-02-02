package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	x := abc(ch)
	fmt.Print(x)
	select {}
}

func abc(ch chan int) int {
	go hello(ch)
	go hello(ch)
	val := <-ch
	return val
}

func hello(ch chan int) {
	ch <- 10
}

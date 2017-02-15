package main

import (
	"fmt"
)

type lol struct {
	a int
	b int
}

func main() {
	a := 0
	item := &lol{
		a: 1,
		b: 2,
	}
	for i := 0; i < 1000; i++ {
		go func() {
			a = a + 1
			item.a = 5
		}()
		go func() {
			//a = a + 2
			item.b = 6
		}()
	}
	fmt.Println(item)
	fmt.Println(a)
}

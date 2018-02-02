package main

import (
	"fmt"
)

func mapTest() {
	x := map[string]interface{}{
		"hello": 1,
	}
	ll := x["ksdfj"]
	fmt.Print(ll)
}

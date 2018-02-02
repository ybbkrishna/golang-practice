package main

import (
	"fmt"
)

func switchStmt() {
	var lol interface{}
	lol = map[string]interface{}{
		"hello": 1,
	}
	// lol = []interface{}{
	// 	1, 2, 3,
	// }
	xx := lol.(map[interface{}]interface{})
	fmt.Println(xx)
	switch lol.(type) {
	case bool:
		fmt.Println("bool")
	case map[interface{}]interface{}:
		fmt.Println("map[interface{}][interface{}]")
	case map[string]interface{}:
		fmt.Println("map[string][interface{}]")
	}
}

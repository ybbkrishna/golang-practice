package main

import (
	"fmt"
	"reflect"
)

func arrayInterface() {
	var x interface{} = [...]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(x))
	lol := reflect.ValueOf(x)
	fmt.Println(lol)
}

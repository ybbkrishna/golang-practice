package main

import "fmt"

func main() {
	s := `package main

import "fmt"

func main() {
	s := %v
	end := "%v"
	fmt.Printf(s, fmt.Sprint(end, s, end), end)
	fmt.Println()
	fmt.Println("}")`
	end := "`"
	fmt.Printf(s, fmt.Sprint(end, s, end), end)
	fmt.Println()
	fmt.Println("}")
}

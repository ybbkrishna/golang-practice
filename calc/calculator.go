/*
Calculator package implements basic reverse polish calculator
expression: 1 2 + 3 *
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ybbkrishna/golang-practice/stack"
)

var reader = bufio.NewReader(os.Stdin)
var s = stack.Stack{
	Capacity: 100,
}

func main() {
	for {
		str, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}
		for _, c := range str {
			switch {
			case c >= '0' && c <= '9':
				token += string(c)
			case c == ' ' && token != "":
				r, _ := strconv.Atoi(token)
				s.Push(r)
				token = ""
			case c == '+':
				a, _ := s.Pop()
				b, _ := s.Pop()
				c, _ := a.(int)
				d, _ := b.(int)
				s.Push(d + c)
			case c == '*':
				a, _ := s.Pop()
				b, _ := s.Pop()
				c, _ := a.(int)
				d, _ := b.(int)
				s.Push(d * c)
			case c == '-':
				a, _ := s.Pop()
				b, _ := s.Pop()
				c, _ := a.(int)
				d, _ := b.(int)
				s.Push(d - c)
			case c == '/':
				a, _ := s.Pop()
				b, _ := s.Pop()
				c, _ := a.(int)
				d, _ := b.(int)
				s.Push(d / c)
			}
		}
		fmt.Println(s.Pop())
	}
}

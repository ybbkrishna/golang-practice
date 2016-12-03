/*
Create a simple stack which can hold a fixed number of ints.
It does not have to grow beyond this limit.
Define
	push – put something on the stack – and
	pop – retrieve something from the stack – functions.
The stack should be a LIFO (last in, first out) stack.
i++ push(k) i     pop() i-- k
0
2. Bonus: Write a String method which converts the stack to a string representation.
This way you can print the stack using: fmt.Printf("My stack %v\n", stack )
The stack in the figure could be represented as: [0:m] [1:l] [2:k]
*/
package main

import (
	"fmt"
	"strings"
)

type stack struct {
	head     *node
	length   int
	capacity int
}

type node struct {
	next  *node
	value interface{}
}

func (s *stack) push(elem interface{}) (bool, error) {
	if s.length == s.capacity {
		panic("Stack Overflow")
	}
	node := node{
		next:  s.head,
		value: elem}
	s.head = &node
	s.length++
	return true, nil
}

func (s *stack) pop() (interface{}, error) {
	if s.length == 0 {
		panic("Empty stack")
	}
	node := s.head
	s.head = node.next
	s.length--
	return node.value, nil
}

func (s stack) Len() int {
	return s.length
}

func (s stack) String() string {
	str := []string{}
	node := s.head
	for i := 0; i < s.length; i++ {
		val := fmt.Sprint(node.value)
		str = append(str, val)
		node = node.next
	}
	return strings.Join(str, " ")
}

func main() {
	s := stack{
		capacity: 5,
	}
	fmt.Println(s)
	s.push("1")
	fmt.Println(s)
	s.push("2")
	s.push("3")
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
	s.push(6)
	val, _ := s.pop()
	_, err := val.(string)
	fmt.Println(err)
}

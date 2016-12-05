/*
	The stack package implements a basic stack with push pop capabilities
*/
package stack

import (
	"fmt"
	"strings"
)

type Stack struct {
	head     *node
	length   int
	Capacity int
}

type node struct {
	next  *node
	value interface{}
}

func (s *Stack) Push(elem interface{}) (bool, error) {
	if s.length == s.Capacity {
		panic("Stack Overflow")
	}
	node := node{
		next:  s.head,
		value: elem}
	s.head = &node
	s.length++
	return true, nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.length == 0 {
		panic("Empty stack")
	}
	node := s.head
	s.head = node.next
	s.length--
	return node.value, nil
}

func (s Stack) Len() int {
	return s.length
}

func (s Stack) String() string {
	str := []string{}
	node := s.head
	for i := 0; i < s.length; i++ {
		val := fmt.Sprint(node.value)
		str = append(str, val)
		node = node.next
	}
	return strings.Join(str, " ")
}

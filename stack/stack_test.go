package stack

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("stackoverflow error expected")
			t.Fail()
		}
	}()
	s := Stack{
		capacity: 1,
	}
	s.Push(1)
	if fmt.Sprint(s) != "1" || s.length != 1 {
		t.Log("stack should contain just one")
		t.Fail()
	}
	s.Push(2)
}

func TestPop(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Log("empty stack error expected")
			t.Fail()
		}
	}()
	s := Stack{
		capacity: 1,
	}
	s.Push(1)
	val, _ := s.Pop()
	if v, ok := val.(int); !ok || v != 1 {
		t.Log("output should be 1")
		t.Fail()
	}
	s.Pop()
}

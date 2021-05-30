package code_02_get_min_stack

import (
	"github.com/golang-collections/collections/stack"
)

type MinStack struct {
	stk *stack.Stack
	minStk *stack.Stack
}

func (s *MinStack) Init() {
	s.stk = stack.New()
	s.minStk = stack.New()
}

func (s *MinStack) Push(x int) {
	s.stk.Push(x)
	if s.minStk.Len() == 0 || x < s.minStk.Peek().(int) {
		s.minStk.Push(x)
	} else {
		x := s.minStk.Peek()
		s.minStk.Push(x)
	}
}

func (s *MinStack) Pop() {
	if s.stk.Len() == 0 {
		panic("the stack is empty")
	}
	s.stk.Pop()
	s.minStk.Pop()
}

func (s *MinStack) GetMinValue() (x int) {
	if s.stk.Len() == 0 {
		panic("the stack is empty")
	}

	x = s.minStk.Peek().(int)
	return x
}

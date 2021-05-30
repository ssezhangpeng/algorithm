package code_03_stack_to_stack

import "github.com/golang-collections/collections/stack"

type TwoStackQueue struct {
	stk1 *stack.Stack
	stk2 *stack.Stack
}

func (s2q *TwoStackQueue) Init() {
	s2q.stk1 = stack.New()
	s2q.stk2 = stack.New()
}

func (s2q *TwoStackQueue) EnQueue(x int) {
	s2q.stk1.Push(x)
}

func (s2q *TwoStackQueue) DeQueue() (x int) {
	if s2q.stk1.Len() == 0 && s2q.stk2.Len() == 0 {
		panic("the stack is empty")
	}

	if s2q.stk2.Len() == 0 {
		for s2q.stk1.Len() != 0 {
			val := s2q.stk1.Pop()
			s2q.stk2.Push(val)
		}
	}

	x = s2q.stk2.Pop().(int)
	return x
}

func (s2q *TwoStackQueue) Peek() (x int) {
	if s2q.stk1.Len() == 0 && s2q.stk2.Len() == 0 {
		panic("the stack is empty")
	}

	if s2q.stk2.Len() == 0 {
		for s2q.stk1.Len() != 0 {
			val := s2q.stk1.Pop()
			s2q.stk2.Push(val)
		}
	}

	x = s2q.stk2.Peek().(int)
	return x
}

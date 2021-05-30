package code_01_array_to_stack_queue

import (
	"fmt"
	"testing"
)

func TestArrayStack(t *testing.T) {
	var stk ArrayStack
	stk.Init(3)

	stk.Push(1)
	stk.Push(2)
	stk.Push(3)

	x := stk.Top()
	fmt.Println(x)
	x = stk.Pop()
	fmt.Println(x)

	x = stk.Top()
	fmt.Println(x)
	x = stk.Pop()
	fmt.Println(x)

	x = stk.Top()
	fmt.Println(x)
	x = stk.Pop()
	fmt.Println(x)

	x = stk.Top()
	fmt.Println(x)
	x = stk.Pop()
	fmt.Println(x)
}

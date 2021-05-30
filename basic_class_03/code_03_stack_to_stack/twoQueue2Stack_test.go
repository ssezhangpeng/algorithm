package code_03_stack_to_stack

import (
	"fmt"
	"testing"
)

func TestTwoQueueStack(t *testing.T) {
	var twoQueueStack TwoQueueStack

	twoQueueStack.Init()

	twoQueueStack.Push(2)
	twoQueueStack.Push(4)
	twoQueueStack.Push(1)

	x := twoQueueStack.Peek()  // 1
	fmt.Println(x)
	x = twoQueueStack.Pop()  // 1
	fmt.Println(x)

	x = twoQueueStack.Peek()  // 4
	fmt.Println(x)
	x = twoQueueStack.Pop()  // 4
	fmt.Println(x)

	x = twoQueueStack.Peek()  // 2
	fmt.Println(x)
	x = twoQueueStack.Pop()  // 2
	fmt.Println(x)
}

package code_03_stack_to_stack

import (
	"fmt"
	"testing"
)

func TestTwoStackQueue (t *testing.T) {
	var twoStackQueue TwoStackQueue

	twoStackQueue.Init()

	twoStackQueue.EnQueue(2)
	twoStackQueue.EnQueue(4)
	twoStackQueue.EnQueue(1)

	x := twoStackQueue.Peek()  // 2
	fmt.Println(x)
	x = twoStackQueue.DeQueue() // 2
	fmt.Println(x)

	x = twoStackQueue.Peek()  // 4
	fmt.Println(x)
	x = twoStackQueue.DeQueue()  // 4
	fmt.Println(x)

	x = twoStackQueue.Peek()  // 1
	fmt.Println(x)
	x = twoStackQueue.DeQueue()  // 1
	fmt.Println(x)

}

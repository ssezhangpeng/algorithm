package code_01_array_to_stack_queue

import (
	"fmt"
	"testing"
)

func TestArrayCircleQueue(t *testing.T) {
	var q ArrayCircleQueue
	q.Init(3)

	q.Push(1)
	q.Push(2)
	q.Push(3)

	x := q.Front()
	fmt.Println(x)  // 1
	x = q.Pop()
	fmt.Println(x)  // 1

	x = q.Front()
	fmt.Println(x)  // 2
	x = q.Pop()
	fmt.Println(x)  // 2

	x = q.Front()
	fmt.Println(x)  // 3
	x = q.Pop()
	fmt.Println(x)  // 3
}

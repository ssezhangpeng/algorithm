package code_02_get_min_stack

import (
	"fmt"
	"testing"
)

func TestGetMinValue(t *testing.T) {
	var minStack MinStack

	minStack.Init()

	minStack.Push(3)
	minStack.Push(2)
	minStack.Push(8)
	minStack.Push(1)

	x := minStack.GetMinValue()
	fmt.Println(x)
}

package code_01_array_to_stack_queue

type ArrayStack struct {
	arr  []int
	len  int
	size int
}

func (stk *ArrayStack) Init(initSize int) {
	stk.arr = make([]int, initSize)
	stk.len = 0
	stk.size = initSize
}

func (stk *ArrayStack) Push(x int) {
	if stk.len == stk.size {
		panic("the stack is full")
	}
	stk.arr[stk.len] = x
	stk.len++
}

func (stk *ArrayStack) Pop() int {
	if stk.len == 0 {
		panic("the stack is empty")
	}
	x := stk.arr[stk.len-1]
	stk.len--
	return x
}

func (stk *ArrayStack) Top() int {
	if stk.len == 0 {
		panic("the stack is empty")
	}

	return stk.arr[stk.len-1]
}

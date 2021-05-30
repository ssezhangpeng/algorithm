package code_01_array_to_stack_queue

type ArrayCircleQueue struct {
	arr []int
	len int
	size int
	first int
	last int
}

func (q *ArrayCircleQueue) Init(initSize int) {
	q.arr = make([]int, initSize)
	q.len = 0
	q.size = initSize
	q.first = 0
	q.last = 0
}

func (q *ArrayCircleQueue) Push(x int) {
	if q.len == q.size {
		panic("the queue is full")
	}
	q.arr[q.first] = x
	if q.first == q.size-1 {
		q.first = 0
	} else {
		q.first++
	}

	q.len++
}

func (q *ArrayCircleQueue) Pop() (x int) {
	if q.len == 0 {
		panic("the queue is empty")
	}

	x = q.arr[q.last]
	if q.last == q.size-1 {
		q.last = 0
	} else {
		q.last++
	}

	q.len--
	return x
}

func (q *ArrayCircleQueue) Front() (x int) {
	if q.len == 0 {
		panic("the queue is empty")
	}

	x = q.arr[q.last]
	return x
}

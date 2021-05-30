package code_03_stack_to_stack

import (
	"github.com/golang-collections/collections/queue"
)

type TwoQueueStack struct {
	q1 *queue.Queue
	q2 *queue.Queue
}

func (q2s *TwoQueueStack) Init() {
	q2s.q1 = queue.New()
	q2s.q2 = queue.New()
}

func (q2s *TwoQueueStack) Push(x int) {
	q2s.q1.Enqueue(x)
}

func (q2s *TwoQueueStack) Pop() (x int) {
	if q2s.q1.Len() == 0 && q2s.q2.Len() == 0{
		panic("the queue is empty")
	}
	for q2s.q1.Len() != 1 {
		val := q2s.q1.Dequeue().(int)
		q2s.q2.Enqueue(val)
	}
	x = q2s.q1.Dequeue().(int)

	q2s.q1, q2s.q2 = q2s.q2, q2s.q1
	return x
}

func (q2s *TwoQueueStack) Peek() (x int) {
	if q2s.q1.Len() == 0 && q2s.q2.Len() == 0{
		panic("the queue is empty")
	}

	for q2s.q1.Len() != 1 {
		val := q2s.q1.Dequeue().(int)
		q2s.q2.Enqueue(val)
	}

	x = q2s.q1.Peek().(int)
	return x
}
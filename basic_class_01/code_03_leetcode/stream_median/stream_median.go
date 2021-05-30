package stream_median

import (
	"container/heap"
)

var (
	minHeap = &IntMinHeap{}  // storage element <= median
	maxHeap = &IntMaxHeap{}  // storage element >= median
)

func init() {
	heap.Init(minHeap)
	heap.Init(maxHeap)
}

func insertElement(num int) {
	if maxHeap.Len() == 0 || maxHeap.Top().(int) > num {
		heap.Push(maxHeap, num)
	} else {
		heap.Push(minHeap, num)
	}

	if maxHeap.Len() - minHeap.Len() > 1 {
		x := heap.Pop(maxHeap)
		heap.Push(minHeap, x)
	} else if minHeap.Len() - maxHeap.Len() > 1 {
		x := heap.Pop(minHeap)
		heap.Push(maxHeap, x)
	}
}

func getMedian() float32 {
	total := minHeap.Len() + maxHeap.Len()
	if total % 2 == 0 {
		return float32(minHeap.Top().(int)+maxHeap.Top().(int)) / 2
	} else {
		if minHeap.Len() > maxHeap.Len() {
			return float32(minHeap.Top().(int))
		}
		return float32(maxHeap.Top().(int))
	}
}

//contain queue 的使用
//contain linklist 的使用



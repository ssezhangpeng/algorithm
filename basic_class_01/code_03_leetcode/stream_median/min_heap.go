package stream_median

//contain heap 的使用
type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]  // if h[i] > h[j]--> big Heap
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x interface{}) {
	// TODO: Push and Pop use pointer receivers because they modify the slice's length, not just contents.
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

// defined by myself
func (h *IntMinHeap) Top() interface{} {
	return (*h)[0]
}

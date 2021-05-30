package code_03_heap_sort

func heapSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	// build big heap
	for i:=0; i<len(arr); i++ {
		heapInsert(arr, i)
	}

	r := len(arr)-1
	arr[0], arr[r] = arr[r], arr[0]

	for r > 0 {
		heapify(arr, 0, r)
		r--
		arr[0], arr[r] = arr[r], arr[0]
	}

}

func heapInsert(arr []int, index int) {
	for arr[index] > arr[(index-1)/2] {  // when index == 0, is a little trick
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index-1)/2
	}
}

func heapify(arr []int, l, r int) {
	lchild := 2 * l +1

	for lchild < r {
		largestIndex := lchild
		if lchild+1 < r {
			if arr[lchild] < arr[lchild+1] {
				largestIndex = lchild + 1
			}
		}
		if arr[largestIndex] <= arr[l] {
			break
		}

		arr[l], arr[largestIndex] = arr[largestIndex], arr[l]
		l = largestIndex
		lchild = 2 * l + 1
	}
}

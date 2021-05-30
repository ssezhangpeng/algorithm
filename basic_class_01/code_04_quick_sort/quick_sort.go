package code_04_quick_sort

func quickSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	quickSortCore(arr, 0, len(arr)-1)
}

func quickSortCore(arr []int, l, r int) {
	if l < r {
		pos := partition(arr, l, r)
		quickSortCore(arr, l, pos[0]-1)
		quickSortCore(arr, pos[1]+1, r)
	}
}

func partition(arr []int, l, r int) (pos []int) {
	pivot := arr[r]
	less := l-1  // point the last element where arr[i] < pivot
	more := r    // point the first element where arr[i] > pivot
	idx := l

	for idx < more {
		if arr[idx] < pivot {
			less++
			arr[less], arr[idx] = arr[idx], arr[less]
			idx++
		} else if arr[idx] > pivot {
			more--
			arr[idx], arr[more] = arr[more], arr[idx]
		} else {
			idx++
		}
	}
	arr[more], arr[r] = arr[r], arr[more]
	return []int{less+1, more}
}


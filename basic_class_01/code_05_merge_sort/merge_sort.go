package code_05_merge_sort

func mergeSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	mergeSortCore(arr, 0, len(arr)-1)
}

func mergeSortCore(arr []int, l, r int) {
	if l == r {
		return
	}

	mid := l + (r-l) >> 1
	mergeSortCore(arr, l, mid)
	mergeSortCore(arr, mid+1, r)
	merge(arr, l, mid, r)
}

func merge(arr []int, l, mid, r int) {
	temp := make([]int, r-l+1)
	p1 := l
	p2 := mid + 1
	idx := 0

	for p1<=mid && p2<=r {
		if arr[p1] < arr[p2] {
			temp[idx] = arr[p1]
			p1++
		} else {
			temp[idx] = arr[p2]
			p2++
		}
		idx++
	}

	for p1 <= mid {
		temp[idx] = arr[p1]
		idx++
		p1++
	}

	for p2 <= r {
		temp[idx] = arr[p2]
		idx++
		p2++
	}

	for i:=0; i<r-l+1; i++ {
		arr[l+i] = temp[i]
	}
}

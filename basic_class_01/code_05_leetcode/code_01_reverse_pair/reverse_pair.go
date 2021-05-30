package code_01_reverse_pair

func reversePair(arr []int) (sum int) {
	if arr == nil || len(arr) < 2 {
		return 0
	}
	return mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) (sum int) {
	if l == r {
		return 0;
	}

	mid := l + (r-l)>>1
	return mergeSort(arr, l, mid) + mergeSort(arr, mid+1, r) + merge(arr, l, mid ,r)
}

func merge(arr []int, l, mid, r int) (mergeSum int) {
	temp := make([]int, r-l+1)
	idx := 0
	p1 := l
	p2 := mid + 1

	for p1<=mid && p2<=r {
		if arr[p1] <= arr[p2] {
			temp[idx] = p1
			p1++
		} else {
			mergeSum += mid - p1 + 1
			temp[idx] = p2
			p2++
		}
		idx++
	}

	for p1 <= mid {
		temp[idx] = p1
		p1++
		idx++
	}

	for p2 <= r {
		temp[idx] = p2
		p2++
		idx++
	}

	for i:=0; i<r-l+1; i++ {
		arr[l+i] = temp[i]
	}

	return mergeSum
}

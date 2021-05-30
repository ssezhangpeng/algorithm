package code_00_partition_array

func partitionArray(arr []int, num int) {
	l := -1
	r := len(arr)
	idx := 0

	for idx < r {
		if arr[idx] <= num {
			l++
			idx++
		} else {
			r--
			arr[idx], arr[r] = arr[r], arr[idx]
		}
	}
	return
}

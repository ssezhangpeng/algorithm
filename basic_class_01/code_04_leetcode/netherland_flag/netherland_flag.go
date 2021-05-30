package netherland_flag

func netherlandFlag(arr []int, num int) {
	l := -1
	r := len(arr)
	idx := 0

	for idx < r {
		if arr[idx] < num {
			l++
			arr[l], arr[idx] = arr[idx], arr[l]
			idx++
		} else if arr[idx] > num {
			r--
			arr[r], arr[idx] = arr[idx], arr[r]
		} else {
			idx++
		}
	}
}

package code_01_insert_sort

func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

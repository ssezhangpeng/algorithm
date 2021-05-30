package code_02_selection_sort

func selectionSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	for i:=0; i<len(arr)-1; i++ {
		minValueIndex := i
		for j:=i; j<len(arr); j++ {
			if arr[j] < arr[minValueIndex] {
				minValueIndex = j
			}
		}
		arr[i], arr[minValueIndex] = arr[minValueIndex], arr[i]
	}
}

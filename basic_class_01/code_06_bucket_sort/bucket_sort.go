package code_06_bucket_sort

import "math"

func bucketSort(arr []int) {
	// only for no-negative value
	if arr == nil || len(arr) < 2 {
		return
	}

	maxVal := math.MinInt8
	for i:=0; i<len(arr); i++ {
		maxVal = int(math.Max(float64(maxVal), float64(arr[i])))
	}

	bucket := make([]int, maxVal+1)
	for i:=0; i<len(arr); i++ {
		bucket[arr[i]]++
	}

	k := 0
	for j:=0; j<len(bucket); j++ {
		for bucket[j] > 0 {
			arr[k] = j
			k++
			bucket[j]--
		}
	}
}

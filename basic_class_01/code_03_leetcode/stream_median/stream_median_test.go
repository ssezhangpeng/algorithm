package stream_median

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func comparator(arr []int) (median float32){
	sort.Ints(arr)
	len := len(arr)
	if len % 2 == 1 {
		return float32(arr[len/2])
	}

	x := float32(arr[len/2]+arr[(len-1)/2]) / 2
	return x
}

func generateRandomArray(maxSize, maxValue int) (arr []int) {
	rand.Seed(time.Now().Unix())
	randomSize := int(float32(maxSize+1) * rand.Float32())
	arr = make([]int, randomSize)
	for i := 0; i < len(arr); i++ {
		arr[i] = int(float32(maxValue+1)*rand.Float32() - float32(maxValue+1)*rand.Float32())
	}
	return arr
}

func isEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func TestMedianSort(t *testing.T) {
	testCnt := 10000
	maxSize := 100
	maxValue := 100
	success := true

	for i := 0; i < testCnt; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := make([]int, len(arr))
		copy(arr1, arr)

		for i:=0; i<len(arr); i++ {
			insertElement(arr[i])
		}
		want := comparator(arr1)
		ans := getMedian()
		if want != ans {
			success = false
			break
		}
	}
	if !success {
		panic("fucking fucked")
	} else {
		fmt.Println("passed")
	}
}

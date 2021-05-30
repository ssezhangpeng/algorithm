package code_00_small_sum

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func comparator(arr []int) (want int) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				want += arr[j]
			}
		}
	}
	return want
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

func TestSmallSum(t *testing.T) {
	testCnt := 10000
	maxSize := 3
	maxValue := 100
	success := true

	for i := 0; i < testCnt; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		arr1 := make([]int, len(arr))
		arr2 := make([]int, len(arr))
		copy(arr1, arr)
		copy(arr2, arr)
		want := comparator(arr1)
		ans := smallSum(arr2)
		if ans != want {
			fmt.Printf("arr: %+v, ans: %d, want: %d\n", arr, ans, want)
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

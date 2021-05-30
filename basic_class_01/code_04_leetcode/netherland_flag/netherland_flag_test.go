package netherland_flag

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func generateRandomArray(maxSize, maxValue int) (arr []int) {
	rand.Seed(time.Now().Unix())
	randomSize := int(float32(maxSize+1) * rand.Float32())
	arr = make([]int, randomSize)
	for i := 0; i < len(arr); i++ {
		arr[i] = int(float32(maxValue+1)*rand.Float32() - float32(maxValue+1)*rand.Float32())
	}
	return arr
}

func judge(arr []int, num int) bool {
	i := 0
	// judge <
	for i<len(arr) {
		if arr[i] >= num {
			break
		}
		i++
	}

	if i == len(arr) {
		return true
	}

	// judge ==
	for i<len(arr){
		if arr[i] != num {
			break
		}
		i++
	}

	if i == len(arr) {
		return true
	}

	// judge >
	for i<len(arr) {
		if arr[i] < num {
			return false
		}
		i++
	}
	return true
}

func TestNetherlandFlag(t *testing.T) {
	testCnt := 100000
	maxSize := 100
	maxValue := 100
	success := true

	for i := 0; i < testCnt; i++ {
		arr := generateRandomArray(maxSize+1, maxValue)
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		netherlandFlag(arr1, arr1[0])
		success = judge(arr1, arr[0])
	}
	if !success {
		panic("fucking fucked")
	} else {
		fmt.Println("passed")
	}
}

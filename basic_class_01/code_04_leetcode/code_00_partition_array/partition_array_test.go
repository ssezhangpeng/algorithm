package code_00_partition_array

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
	for i<len(arr) {
		if arr[i] > num {
			break
		}
		i++
	}

	if i == len(arr) {
		return true
	}

	for j:=i; j<len(arr); j++ {
		if arr[j] < num {
			return false
		}
	}
	return true
}

func TestPartitionArray(t *testing.T) {
	testCnt := 10000
	maxSize := 100
	maxValue := 100
	success := true

	for i := 0; i < testCnt; i++ {
		arr := generateRandomArray(maxSize+1, maxValue)
		partitionArray(arr, arr[0])
		success = judge(arr, arr[0])
	}
	if !success {
		panic("fucking fucked")
	} else {
		fmt.Println("passed")
	}
}
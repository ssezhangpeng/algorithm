package code_06_leetcode

import (
	"math"
)

func maxGap(arr []int) int {
	if arr == nil || len(arr) < 2 {
		return 0
	}

	len := len(arr)
	minVal := math.MaxInt8
	maxVal := math.MinInt8

	for i:=0; i<len; i++ {
		minVal = int(math.Min(float64(minVal), float64(arr[i])))
		maxVal = int(math.Max(float64(maxVal), float64(arr[i])))
	}

	// bucket use (len/(maxVal-minVal)) divide by zero
	if minVal == maxVal {
		return 0
	}

	mins := make([]int, len+1)
	maxs := make([]int, len+1)
	hasNum := make([]bool, len+1)

	for i:=0; i<len; i++ {
		bid := bucket(arr[i], len, minVal, maxVal)
		if hasNum[bid] {
			mins[bid] = int(math.Min(float64(mins[bid]), float64(arr[i])))
			maxs[bid] = int(math.Max(float64(maxs[bid]), float64(arr[i])))
		} else {
			mins[bid] = arr[i]
			maxs[bid] = arr[i]
			hasNum[bid] = true
		}
	}

	ans := 0
	lastMax := maxs[0]
	for i:=1; i<=len; i++ {
		if hasNum[i] {
			ans = int(math.Max(float64(ans), float64(mins[i]-lastMax)))
			lastMax = maxs[i]
		}
	}

	return ans
}

func bucket(num, len, minVal, maxVal int) int {
	return int((num-minVal)*len/(maxVal-minVal))
}

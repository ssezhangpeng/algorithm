package code_05_rotate_matrix

import (
	"fmt"
	"testing"
)

func printMatrix(matrix [][]int) {
	for i:=0; i<len(matrix); i++ {
		for j:=0; j<len(matrix[0]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func TestRotateMatrix(t *testing.T) {
	matrix := [][]int {
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	}

	RotateMatrix(matrix)
	printMatrix(matrix)
}

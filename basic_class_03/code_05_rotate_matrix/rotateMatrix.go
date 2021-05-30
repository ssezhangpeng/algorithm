package code_05_rotate_matrix

func RotateMatrix(matrix [][]int) {
	tR := 0
	tC := 0
	dR := len(matrix) - 1
	dC := len(matrix[0]) - 1

	for tR < dR {
		RotateMatrixCore(matrix, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

func RotateMatrixCore(matrix [][]int, tR, tC, dR, dC int) {
	times := dC - tC
	temp := 0

	for i:=0; i<times; i++ {
		temp = matrix[tR][tC+i]
		matrix[tR][tC+i] = matrix[dR-i][tC]
		matrix[dR-i][tC] = matrix[dR][dC-i]
		matrix[dR][dC-i] = matrix[tR+i][dC]
		matrix[tR+i][dC] = temp
	}
}

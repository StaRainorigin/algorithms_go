package solution

import "slices"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := range n / 2 {
		for j := range n - 1 - 2*i {
			matrix[0+i][0+i+j], matrix[0+i+j][n-1-i], matrix[n-1-i][n-1-i-j], matrix[n-1-i-j][0+i] = 
			matrix[n-1-i-j][0+i], matrix[0+i][0+i+j], matrix[0+i+j][n-1-i], matrix[n-1-i][n-1-i-j]
		}
	}
}

func rotate1(matrix [][]int) {
	n := len(matrix)
	for i, row := range matrix {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		slices.Reverse(row)
	}
}
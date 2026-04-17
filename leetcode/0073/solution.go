package solution

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	firstRowZero := 1
	
	for j := range n {
		if matrix[0][j] == 0 {
			firstRowZero = 0
			break
		}
	}
	for i := range m {
		if matrix[i][0] == 0 {
			matrix[0][0] = 0
			break
		}
	}
	
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}
	
	if matrix[0][0] == 0 {
		for i := 1; i < m; i++ {
			matrix[i][0] = 0
		}
	}
	if firstRowZero == 0 {
		for j := range n {
			matrix[0][j] = 0
		}
	}
}

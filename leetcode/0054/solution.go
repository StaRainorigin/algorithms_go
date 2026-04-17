package solution

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])
	if row <= 0 {
		return []int{}
	}
	
	check := make([][]bool, row)
	for i := range check {
		check[i] = make([]bool, col)
	}
	
	operation := 0
	dx := [4]int{0, 1, 0, -1}
	dy := [4]int{1, 0, -1, 0}
	
	nextStep := func(i, j int) (int, int) {
		ni, nj := i+dx[operation], j+dy[operation]
		if ni < 0 || ni >= row || nj < 0 || nj >= col || check[ni][nj] {
			operation = (operation + 1) % 4
			ni, nj = i+dx[operation], j+dy[operation]
		}
		return ni, nj
	}
	
	ans := make([]int, row*col)
	i, j := 0, 0
	for k := range row*col{
		ans[k] = matrix[i][j]
		check[i][j] = true
		i, j = nextStep(i, j)
	}
	
	return ans
}

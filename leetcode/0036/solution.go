package solution

func isValidSudoku(board [][]byte) bool {
	checkRow, checkCol, checkGrid := [9][9]bool{}, [9][9]bool{}, [3][3][9]bool{}
	for i := range 9 {
		for j := range 9 {
			x := board[i][j]
			if x == '.' {
				continue
			}
			x = x - '1'
			if checkRow[i][x] || checkCol[x][j] || checkGrid[i/3][j/3][x] {
				return false
			}
			checkRow[i][x] = true
			checkCol[x][j] = true
			checkGrid[i/3][j/3][x] = true
		}
	}
	return true
}

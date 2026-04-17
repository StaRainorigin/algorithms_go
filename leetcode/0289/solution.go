package solution

func gameOfLife(board [][]int) {
	m := len(board)
	if m <= 0 {
		return
	}
	n := len(board[0])
	updateLife := make([][]int, m)
	for i := range m {
		updateLife[i] = make([]int, n)
	}
	
	check := func (i, j int) int {
		count := 0 
		hasTop, hasLeft, hasBottom, HasRight := i > 0, j > 0, i < m - 1, j < n - 1
		if hasTop {
			count += board[i-1][j]
			if hasLeft {
				count += board[i-1][j-1]
			}
			if HasRight {
				count += board[i-1][j+1]
			}
		}
		if hasBottom {
			count += board[i+1][j]
			if hasLeft {
				count += board[i+1][j-1]
			}
			if HasRight {
				count += board[i+1][j+1]
			}
		}
		if hasLeft {
			count += board[i][j-1]
		}
		if HasRight {
			count += board[i][j+1]
		}
		return count
	}
	
	for i := range m {
		for j := range n {
			alive := check(i, j) 
			cur := board[i][j]
			if cur == 1 && (alive < 2 || alive > 3) {
				updateLife[i][j] = 0
				continue
			}
			if cur == 0 && alive == 3 {
				updateLife[i][j] = 1
				continue
			}
			updateLife[i][j] = cur
		}
	}
	
	for i := range m {
	    for j := range n {
	        board[i][j] = updateLife[i][j]
	    }
	}
}


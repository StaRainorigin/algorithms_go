package solution

var DIRS = [][]int{
	[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1},
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] != 'O' {
			return
		}
		board[i][j] = '#'

		for _, dir := range DIRS {
			dfs(i + dir[0], j + dir[1])
		}
	}

	for i := range m {
		dfs(i, 0)
		dfs(i, n - 1)
	}
	for j := range n {
		dfs(0, j)
		dfs(m - 1, j)
	}

	for i := range m {
		for j := range n {
			if board[i][j] == '#' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

// func solve(board [][]byte) {
// 	m, n := len(board), len(board[0])
// 	vis := make([][]bool, m)
// 	for i := range vis {
// 		vis[i] = make([]bool, n)
// 	}

// 	var dfs func(int, int, int) bool
// 	dfs = func(i, j, op int) bool {
// 		if i < 0 || j < 0 || i >= m || j >= n || board[i][j] == 'X' || (op == 0 && vis[i][j]) {
// 			return true
// 		}

// 		if op == 0 { // op == 0 是检索是否贴边
// 			vis[i][j] = true
// 			if i == 0 || j == 0 || i == m-1 || j == n-1 {
// 				return false
// 			}

// 			res := true
// 			for _, dir := range DIRS {
// 				res = dfs(i + dir[0], j + dir[1], 0) && res
// 			}
// 			return res

// 		} else { // op == 1 是确定不贴边之后，第二次遍历把'X'改成'O'
// 			board[i][j] = 'X'
// 			for _, dir := range DIRS {
// 				dfs(i + dir[0], j + dir[1], 1)
// 			}
// 		}

// 		return true
// 	}

// 	for i, row := range board {
// 		for j := range row {
// 			if dfs(i, j, 0) {
// 				dfs(i, j, 1)
// 			}
// 		}
// 	}
// }

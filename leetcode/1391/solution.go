package solution

import "slices"

var dirs = [7][2][2]int{
	{},
	{{0, -1}, {0, 1}},  // 站在街道 1，可以往左或者往右
	{{-1, 0}, {1, 0}},  // 站在街道 2，可以往上或者往下
	{{0, -1}, {1, 0}},  // 站在街道 3，可以往左或者往下
	{{0, 1}, {1, 0}},   // 站在街道 4，可以往右或者往下
	{{0, -1}, {-1, 0}}, // 站在街道 5，可以往左或者往上
	{{0, 1}, {-1, 0}},  // 站在街道 6，可以往右或者往上
}

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	if m == 0 {
		return true
	}
	n := len(grid[0])
	vis := make([][]bool, m) 
	for i := range m {
		vis[i] = make([]bool, n)
	}
	
	var dfs func(int, int) bool
	dfs = func(x, y int) bool {
		if x == m - 1 && y == n - 1 {
			return true
		}
		vis[x][y] = true
		for _, d := range dirs[grid[x][y]] {
			i, j := x + d[0], y + d[1]
			if 0 <= i && i < m && 0 <= j && j < n && !vis[i][j] && 
			slices.Contains(dirs[grid[i][j]][:], [2]int{-d[0], -d[1]}) && dfs(i, j) {
				return true
			}
		}
		return false
	}
	
	return dfs(0, 0)
}

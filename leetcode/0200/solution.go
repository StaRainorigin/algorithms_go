package solution

var DIRS [][]int = [][]int{
	[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1},
} 

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	
	var dfs func(int, int) 
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= m || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		for _, dir := range DIRS {
			dfs(i + dir[0], j + dir[1])
		}
	}
	
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '1' {
				ans++
			}
			dfs(i, j)
		}
	}

	return 
}

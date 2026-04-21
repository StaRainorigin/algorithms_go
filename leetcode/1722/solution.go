package solution

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) (ans int) {
	n := len(source)
	g := make([][]int, n)
	for _, e := range allowedSwaps {
		i, j := e[0], e[1]
		g[i] = append(g[i], j) 
		g[j] = append(g[j], i)
	}

	vis := make([]bool, n)
	diff := map[int]int{}

	var dfs func(int)
	dfs = func(x int) {
		vis[x] = true 
		
		diff[source[x]]++
		diff[target[x]]--
		for _, y := range g[x] {
			if !vis[y] {
				dfs(y)
			}
		}
	}

	for x, b := range vis {
		if !b {
			clear(diff)
			dfs(x)
			for _, c := range diff {
				ans += abs(c)
			}
		}
	}
	return ans / 2 
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

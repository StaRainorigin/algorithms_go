package solution

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	q := [][2]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				dist[i][j] = 0
				q = append(q, [2]int{i, j})
			}
		}
	}

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]
		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[r][c] + 1
				q = append(q, [2]int{nr, nc})
			}
		}
	}

	check := func(limit int) bool {
		if dist[0][0] < limit {
			return false
		}
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[0][0] = true
		q := [][2]int{{0, 0}}
		for len(q) > 0 {
			r, c := q[0][0], q[0][1]
			q = q[1:]
			if r == n-1 && c == n-1 {
				return true
			}
			for _, d := range dirs {
				nr, nc := r+d[0], c+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n && !vis[nr][nc] && dist[nr][nc] >= limit {
					vis[nr][nc] = true
					q = append(q, [2]int{nr, nc})
				}
			}
		}
		return false
	}

	lo, hi := 0, 0
	for i := range dist {
		for j := range dist[i] {
			hi = max(hi, dist[i][j])
		}
	}
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if check(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

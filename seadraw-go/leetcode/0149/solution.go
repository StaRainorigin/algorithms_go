package solution

import "math"

func maxPoints(points [][]int) int {
	ans := 0
	for i, p1 := range points {
		x1, y1 := p1[0], p1[1]
		cnt := map[float64]int{}
		for _, p2 := range points[i+1:] {
			x2, y2 := p2[0], p2[1]
			k := math.MaxFloat64
			if x1 != x2 {
				k = float64(y1 - y2) / float64(x1 - x2)
			}
			cnt[k]++
			ans = max(ans, cnt[k])
		}
	}

	return ans + 1
}

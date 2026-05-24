package solution

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	for _, cur := range intervals {
		m := len(ans)
		if m > 0 && ans[m-1][1] >= cur[0] {
			ans[m-1][1] = max(ans[m-1][1], cur[1])
		} else {
			ans = append(ans, cur)
		}
	}
	return ans
}

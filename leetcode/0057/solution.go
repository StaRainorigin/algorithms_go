package solution

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	m := len(intervals)
	i := 0
	for i < m && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}
	for i < m && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	ans = append(ans, newInterval)
	for i < m {
		ans = append(ans, intervals[i])
		i++
	}
	return ans
}

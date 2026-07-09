package solution

func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	roads := make([]int, n)
	for i := 1; i < len(nums); i++{
		if nums[i] - nums[i-1] > maxDiff {
			roads[i] = i
		} else {
			roads[i] = roads[i-1]
		}
	}
	
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = roads[q[0]] == roads[q[1]]
	}

	return ans
}

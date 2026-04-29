package solution

import "slices"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minOperations(grid [][]int, x int) int {
	m := len(grid)
	if m == 0 {
		return -1
	}
	n := len(grid[0])
	nums := make([]int, m * n)
	idx := 0
	for _, row := range grid {
		for _, num := range row {
			nums[idx] = num
			idx++
		}
	}
	slices.Sort(nums)
	mid := nums[len(nums)/2]
	res := 0
	for _, num := range nums {
		cur := abs(num - mid)
		if cur % x != 0 {
			return -1
		}
		res += cur / x
	}
	return res
}

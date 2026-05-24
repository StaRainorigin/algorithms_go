package solution

import (
	"slices"
	"sort"
)

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	slices.Sort(nums)
	sum := make([]int, n + 1)
	for i, x := range nums {
		sum[i+1] = sum[i] + x
	}
	
	ans := make([]int64, len(queries))
	for i, q := range queries {
		j := sort.SearchInts(nums, q)
		left := q * j - sum[j]
		right := sum[n] - sum[j] - q * (n - j)
		ans[i] = int64(left + right)
	}
	return ans
}

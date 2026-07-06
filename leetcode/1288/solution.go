package solution

import (
	"cmp"
	"slices"
)

func removeCoveredIntervals(intervals [][]int) int {
	ans := 0
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	maxRight := 0
	for _, x := range intervals {
		if x[1] > maxRight {
			maxRight = x[1]
			ans++
		}
	}
	
	return ans
}

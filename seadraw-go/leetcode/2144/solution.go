package solution

import "slices"

func minimumCost(cost []int) int {
	slices.Sort(cost)
	ans := 0
	for i := len(cost) - 1; i >= 0; i -= 3 {
		if i - 1 < 0 {
			ans += cost[i]
		} else {
			ans += cost[i] + cost[i-1]
		}
	}
	return ans
}

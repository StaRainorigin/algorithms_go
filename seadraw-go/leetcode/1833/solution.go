package solution

import "slices"

func maxIceCream(costs []int, coins int) int {
	mx := slices.Max(costs)
	cnt := make([]int, mx + 1)
	for _, cost := range costs {
		if cost <= coins {
			cnt[cost]++
		}
	}
	ans := 0
	for i, x := range cnt {
		for range x {
			if coins < i {
				return ans
			} else {
				coins-=i
				ans++
			}
		}
	}
	return ans
}

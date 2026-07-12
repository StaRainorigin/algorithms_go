package solution

import "slices"

func arrayRankTransform(arr []int) []int {
	sortedArr := slices.Clone(arr)
	slices.Sort(sortedArr)
	sortedArr = slices.Compact(sortedArr)

	mp := map[int]int{}
	for i, x := range sortedArr {
		mp[x] = i + 1
	}

	ans := make([]int, len(arr))
	for i, x := range arr {
		ans[i] = mp[x]
	}

	return ans
}
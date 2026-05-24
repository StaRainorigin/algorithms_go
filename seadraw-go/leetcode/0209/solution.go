package solution

import "math"

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	i := 0
	curSum, res := 0, math.MaxInt
	for j := range n {
		curSum += nums[j]
		for curSum >= target {
			res = min(res, j - i + 1)
			curSum -= nums[i]
			i++
		}
	}
	if res == math.MaxInt {
		res = 0
	}
	return res
}

package solution

func leftRightDifference(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ls, rs := 0, 0
	for _, x := range nums {
		rs += x
	}
	for i := range ans {
		rs -= nums[i]
		if ls < rs {
			ans[i] = rs - ls
		} else {
			ans[i] = ls - rs
		}
		ls += nums[i]
	}
	return ans
}

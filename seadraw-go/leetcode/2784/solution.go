package solution

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGood(nums []int) bool {
	n := len(nums) - 1
	cnt_n := 0
	for _, x := range nums {
		x := abs(x)

		if (x > n) || (x == n && cnt_n > 1) || (x < n && nums[x] < 0) {
			return false
		}

		if x < n {
			nums[x] = -nums[x]
		} else {
			cnt_n++
		}
	}
	return true
}

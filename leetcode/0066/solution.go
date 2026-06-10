package solution

func plusOne(digits []int) []int {
	ans := make([]int, len(digits)+1)
	x := 1
	for i := len(digits) - 1; i >= 0; i-- {
		ans[i+1] = (digits[i]+x) % 10
		x = (digits[i]+x) / 10
	}

	if x == 1 {
		ans[0] = 1
		return ans
	}
	return ans[1:]
}

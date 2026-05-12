package solution

import "slices"

func separateDigits(nums []int) []int {
	ans := []int{}
	for _, x := range slices.Backward(nums) {
		for ; x > 0; x /= 10 {
			ans = append(ans, x % 10)
		}
	}
	slices.Reverse(ans)
	return ans
}

package solution

import "math"

func trans(x int) int {
	y := 0
	for ; x > 0; x /= 10 {
		y += x % 10
	}
	return y
}

func minElement(nums []int) int {
	ans := math.MaxInt
	for _, x := range nums {
		y := trans(x)
		ans = min(ans, y)
	}
	return ans
}
